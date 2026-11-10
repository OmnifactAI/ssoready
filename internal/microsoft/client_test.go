package microsoft

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// newTestClient returns a client whose token exchange is answered with an ID
// token carrying claims. The signature is not checked by ExchangeToken: the
// token comes straight from Microsoft's token endpoint over TLS.
func newTestClient(t *testing.T, status int, claims jwt.MapClaims) *Client {
	t.Helper()

	idToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("test"))
	require.NoError(t, err)
	body, err := json.Marshal(tokenResponse{IDToken: idToken})
	require.NoError(t, err)

	return &Client{
		MicrosoftOAuthClientID:     "client-id",
		MicrosoftOAuthClientSecret: "client-secret",
		MicrosoftOAuthRedirectURI:  "https://app.example.com/callback",
		HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			reqBody, err := io.ReadAll(r.Body)
			require.NoError(t, err)
			form, err := url.ParseQuery(string(reqBody))
			require.NoError(t, err)
			assert.Equal(t, "authorization_code", form.Get("grant_type"))
			assert.Equal(t, "the-code", form.Get("code"))
			assert.Equal(t, "client-id", form.Get("client_id"))
			return &http.Response{StatusCode: status, Body: io.NopCloser(strings.NewReader(string(body)))}, nil
		})},
	}
}

func TestExchangeTokenReturnsProfile(t *testing.T) {
	client := newTestClient(t, http.StatusOK, jwt.MapClaims{
		"name": "Jane Doe", "email": "jane@example.com", "tid": "tenant-123",
	})

	profile, err := client.ExchangeToken(context.Background(), "the-code")

	require.NoError(t, err)
	assert.Equal(t, &Profile{TenantID: "tenant-123", Name: "Jane Doe", Email: "jane@example.com"}, profile)
}

func TestExchangeTokenDropsPersonalAccountTenant(t *testing.T) {
	client := newTestClient(t, http.StatusOK, jwt.MapClaims{
		"name": "Jane Doe", "email": "jane@outlook.com", "tid": "9188040d-6c67-4c5b-b112-36a304b66dad",
	})

	profile, err := client.ExchangeToken(context.Background(), "the-code")

	require.NoError(t, err)
	assert.Empty(t, profile.TenantID)
}

func TestExchangeTokenRejectsMissingEmail(t *testing.T) {
	client := newTestClient(t, http.StatusOK, jwt.MapClaims{"name": "Jane Doe", "tid": "tenant-123"})

	_, err := client.ExchangeToken(context.Background(), "the-code")

	assert.ErrorContains(t, err, "missing name or email")
}

func TestExchangeTokenRejectsErrorResponse(t *testing.T) {
	client := newTestClient(t, http.StatusBadRequest, jwt.MapClaims{})

	_, err := client.ExchangeToken(context.Background(), "the-code")

	assert.ErrorContains(t, err, "bad response status code")
}
