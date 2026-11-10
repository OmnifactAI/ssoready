package store_test

// Runs the SAML login flow through the store against a real Postgres database:
// API redirect URL -> auth service init -> assertion -> access code redemption.
// It covers the database layer (pgx, sqlc queries, migrations) that unit tests
// cannot reach.
//
// Set SSOREADY_TEST_DATABASE_URL to an empty, disposable database to run it,
// e.g. postgres://postgres:password@localhost:5432/postgres?sslmode=disable.
// The test applies all migrations to it. Without the variable it is skipped.

import (
	"context"
	"errors"
	"net/url"
	"os"
	"strings"
	"testing"

	"connectrpc.com/connect"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ssoready/ssoready/internal/authn"
	ssoreadyv1 "github.com/ssoready/ssoready/internal/gen/ssoready/v1"
	"github.com/ssoready/ssoready/internal/pagetoken"
	"github.com/ssoready/ssoready/internal/store"
	"github.com/ssoready/ssoready/internal/store/idformat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type fixture struct {
	appOrgID   uuid.UUID
	envID      uuid.UUID
	orgID      uuid.UUID
	samlConnID uuid.UUID
}

func newTestStore(t *testing.T) (*store.Store, *fixture) {
	t.Helper()

	dbURL := os.Getenv("SSOREADY_TEST_DATABASE_URL")
	if dbURL == "" {
		t.Skip("SSOREADY_TEST_DATABASE_URL is not set")
	}

	m, err := migrate.New("file://../../cmd/migrate/migrations", dbURL)
	require.NoError(t, err)
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		require.NoError(t, err)
	}

	ctx := context.Background()
	db, err := pgxpool.New(ctx, dbURL)
	require.NoError(t, err)
	t.Cleanup(db.Close)

	f := &fixture{appOrgID: uuid.New(), envID: uuid.New(), orgID: uuid.New(), samlConnID: uuid.New()}
	for _, stmt := range []struct {
		sql  string
		args []any
	}{
		{`insert into app_organizations (id) values ($1)`, []any{f.appOrgID}},
		{`insert into environments (id, app_organization_id, redirect_url, display_name) values ($1, $2, 'https://app.example.com/sso/callback', 'test')`, []any{f.envID, f.appOrgID}},
		{`insert into organizations (id, environment_id, external_id) values ($1, $2, $3)`, []any{f.orgID, f.envID, "ext-" + f.orgID.String()}},
		{`insert into organization_domains (id, organization_id, domain) values ($1, $2, 'example.com')`, []any{uuid.New(), f.orgID}},
		{`insert into saml_connections (id, organization_id, is_primary, sp_entity_id, sp_acs_url, idp_entity_id, idp_redirect_url, idp_x509_certificate)
		  values ($1, $2, true, 'https://auth.example.com/sp', 'https://auth.example.com/acs', 'https://idp.example.com', 'https://idp.example.com/sso', $3)`,
			[]any{f.samlConnID, f.orgID, []byte("not-a-real-certificate")}},
	} {
		_, err := db.Exec(ctx, stmt.sql, stmt.args...)
		require.NoError(t, err, stmt.sql)
	}

	var stateKey [32]byte
	copy(stateKey[:], "0123456789abcdef0123456789abcdef")

	return store.New(store.NewStoreParams{
		DB:                  db,
		PageEncoder:         pagetoken.Encoder{Secret: stateKey},
		DefaultAuthURL:      "https://auth.example.com",
		SAMLStateSigningKey: stateKey,
	}), f
}

func (f *fixture) apiKeyContext() context.Context {
	return authn.NewContext(context.Background(), authn.ContextData{APIKey: &authn.APIKeyData{
		AppOrgID: f.appOrgID,
		EnvID:    idformat.Environment.Format(f.envID),
		APIKeyID: "test-api-key",
	}})
}

func TestSAMLLoginFlow(t *testing.T) {
	s, f := newTestStore(t)
	ctx := f.apiKeyContext()
	samlConnID := idformat.SAMLConnection.Format(f.samlConnID)

	// The app asks for a redirect URL for the organization.
	redirectRes, err := s.GetSAMLRedirectURL(ctx, &ssoreadyv1.GetSAMLRedirectURLRequest{
		OrganizationExternalId: "ext-" + f.orgID.String(),
		State:                  "app-state",
	})
	require.NoError(t, err)

	redirectURL, err := url.Parse(redirectRes.RedirectUrl)
	require.NoError(t, err)
	assert.Equal(t, "auth.example.com", redirectURL.Host)
	assert.True(t, strings.HasSuffix(redirectURL.Path, "/v1/saml/"+samlConnID+"/init"), redirectURL.Path)
	state := redirectURL.Query().Get("state")

	// The auth service starts the flow with the signed state.
	initRes, err := s.AuthGetInitData(context.Background(), &store.AuthGetInitDataRequest{State: state, SAMLConnectionID: samlConnID})
	require.NoError(t, err)
	assert.Equal(t, "https://idp.example.com/sso", initRes.IDPRedirectURL)

	require.NoError(t, s.AuthUpsertInitiateData(context.Background(), &store.AuthUpsertInitiateDataRequest{
		State:           state,
		InitiateRequest: "<samlp:AuthnRequest/>",
	}))

	// The auth service validates the assertion against the connection's data.
	validateRes, err := s.AuthGetValidateData(context.Background(), &store.AuthGetValidateDataRequest{SAMLConnectionID: samlConnID})
	require.NoError(t, err)
	assert.Equal(t, "https://idp.example.com", validateRes.IDPEntityID)
	assert.Equal(t, []string{"example.com"}, validateRes.OrganizationDomains)

	assertionID := "assertion-" + uuid.NewString()
	assertionRes, err := s.AuthUpsertReceiveAssertionData(context.Background(), &store.AuthUpsertSAMLLoginEventRequest{
		SAMLConnectionID:     samlConnID,
		SAMLAssertionID:      &assertionID,
		SAMLFlowID:           initRes.RequestID,
		Email:                "jane@example.com",
		SubjectIDPAttributes: map[string]string{"role": "admin"},
		SAMLAssertion:        "<saml:Assertion/>",
	})
	require.NoError(t, err)
	require.NotEmpty(t, assertionRes.Token)

	// The app redeems the access code for the user's identity.
	redeemRes, err := s.RedeemSAMLAccessCode(ctx, &ssoreadyv1.RedeemSAMLAccessCodeRequest{SamlAccessCode: assertionRes.Token})
	require.NoError(t, err)
	assert.Equal(t, "jane@example.com", redeemRes.Email)
	assert.Equal(t, map[string]string{"role": "admin"}, redeemRes.Attributes)
	assert.Equal(t, "app-state", redeemRes.State)
	assert.Equal(t, idformat.Organization.Format(f.orgID), redeemRes.OrganizationId)

	// An access code works only once.
	_, err = s.RedeemSAMLAccessCode(ctx, &ssoreadyv1.RedeemSAMLAccessCodeRequest{SamlAccessCode: assertionRes.Token})
	assert.Equal(t, connect.CodeNotFound, connect.CodeOf(err))

	// A replayed assertion is rejected.
	_, err = s.AuthUpsertReceiveAssertionData(context.Background(), &store.AuthUpsertSAMLLoginEventRequest{
		SAMLConnectionID: samlConnID,
		SAMLAssertionID:  &assertionID,
		Email:            "jane@example.com",
		SAMLAssertion:    "<saml:Assertion/>",
	})
	assert.ErrorIs(t, err, store.ErrDuplicateAssertionID)
}

func TestSAMLLoginRejectsTamperedState(t *testing.T) {
	s, f := newTestStore(t)

	redirectRes, err := s.GetSAMLRedirectURL(f.apiKeyContext(), &ssoreadyv1.GetSAMLRedirectURLRequest{
		SamlConnectionId: idformat.SAMLConnection.Format(f.samlConnID),
	})
	require.NoError(t, err)
	redirectURL, err := url.Parse(redirectRes.RedirectUrl)
	require.NoError(t, err)
	payload, _, _ := strings.Cut(redirectURL.Query().Get("state"), ".")

	_, err = s.AuthGetInitData(context.Background(), &store.AuthGetInitDataRequest{
		State:            payload + ".AAAA",
		SAMLConnectionID: idformat.SAMLConnection.Format(f.samlConnID),
	})

	var badState *store.AuthGetInitDataBadStateError
	assert.ErrorAs(t, err, &badState)
}

func TestSAMLRedirectURLIsScopedToEnvironment(t *testing.T) {
	s, f := newTestStore(t)
	_, other := newTestStore(t)

	_, err := s.GetSAMLRedirectURL(other.apiKeyContext(), &ssoreadyv1.GetSAMLRedirectURLRequest{
		SamlConnectionId: idformat.SAMLConnection.Format(f.samlConnID),
	})

	assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
}
