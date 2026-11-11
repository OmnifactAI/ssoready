package google

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testClientID = "test-client-id"

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

// newTestClient returns a client whose HTTP calls to Google's JWKS endpoint are
// answered with the public half of key.
func newTestClient(t *testing.T, key *rsa.PrivateKey) *Client {
	t.Helper()

	jwks, err := json.Marshal(jose.JSONWebKeySet{Keys: []jose.JSONWebKey{{
		Key: &key.PublicKey, KeyID: "test-key", Algorithm: string(jose.RS256), Use: "sig",
	}}})
	require.NoError(t, err)

	return &Client{
		GoogleOAuthClientID: testClientID,
		HTTPClient: &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			require.Equal(t, googleOAuthJWKSURL, r.URL.String())
			return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(string(jwks)))}, nil
		})},
	}
}

func newKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	return key
}

func signCredential(t *testing.T, key *rsa.PrivateKey, claims googleCredentialClaims) string {
	t.Helper()
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: key},
		(&jose.SignerOptions{}).WithType("JWT").WithHeader("kid", "test-key"),
	)
	require.NoError(t, err)

	token, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	require.NoError(t, err)
	return token
}

func validClaims() googleCredentialClaims {
	now := time.Now()
	return googleCredentialClaims{
		Claims: jwt.Claims{
			Issuer:   "https://accounts.google.com",
			Audience: jwt.Audience{testClientID},
			IssuedAt: jwt.NewNumericDate(now),
			Expiry:   jwt.NewNumericDate(now.Add(time.Hour)),
		},
		HostedDomain: "example.com",
		Email:        "jane@example.com",
		Name:         "Jane Doe",
	}
}

func TestParseCredentialAcceptsValidToken(t *testing.T) {
	key := newKey(t)

	res, err := newTestClient(t, key).ParseCredential(context.Background(), &ParseCredentialRequest{
		Credential: signCredential(t, key, validClaims()),
	})

	require.NoError(t, err)
	assert.Equal(t, &ParseCredentialResponse{HostedDomain: "example.com", Email: "jane@example.com", Name: "Jane Doe"}, res)
}

func TestParseCredentialRejectsInvalidTokens(t *testing.T) {
	key := newKey(t)

	tests := []struct {
		name    string
		signKey *rsa.PrivateKey
		modify  func(*googleCredentialClaims)
	}{
		{name: "signed by another key", signKey: newKey(t), modify: func(*googleCredentialClaims) {}},
		{name: "wrong audience", signKey: key, modify: func(c *googleCredentialClaims) { c.Audience = jwt.Audience{"other-client"} }},
		{name: "wrong issuer", signKey: key, modify: func(c *googleCredentialClaims) { c.Issuer = "https://evil.example.com" }},
		{name: "expired", signKey: key, modify: func(c *googleCredentialClaims) {
			c.Expiry = jwt.NewNumericDate(time.Now().Add(-time.Hour))
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims := validClaims()
			tt.modify(&claims)

			_, err := newTestClient(t, key).ParseCredential(context.Background(), &ParseCredentialRequest{
				Credential: signCredential(t, tt.signKey, claims),
			})

			assert.Error(t, err)
		})
	}
}
