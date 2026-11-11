package authservice

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newRSAKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	return key
}

// signIDToken signs claims the same way oauthToken does.
func signIDToken(t *testing.T, alg jose.SignatureAlgorithm, key any, subject string) string {
	t.Helper()
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: alg, Key: key}, (&jose.SignerOptions{}).WithType("JWT"))
	require.NoError(t, err)

	now := time.Now()
	token, err := jwt.Signed(signer).Claims(idTokenClaims{Claims: jwt.Claims{
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(time.Hour)),
		Subject:  subject,
	}}).Serialize()
	require.NoError(t, err)
	return token
}

func callUserinfo(s *Service, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/v1/oauth/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	s.oauthUserinfo(rec, req)
	return rec
}

func TestOAuthUserinfoReturnsSubjectOfValidToken(t *testing.T) {
	key := newRSAKey(t)
	s := &Service{OAuthIDTokenPrivateKey: key}

	rec := callUserinfo(s, signIDToken(t, jose.RS256, key, "jane@example.com"))

	require.Equal(t, http.StatusOK, rec.Code)
	var body struct {
		Sub   string `json:"sub"`
		Email string `json:"email"`
	}
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &body))
	assert.Equal(t, "jane@example.com", body.Sub)
	assert.Equal(t, "jane@example.com", body.Email)
}

func TestOAuthUserinfoRejectsInvalidTokens(t *testing.T) {
	key := newRSAKey(t)
	s := &Service{OAuthIDTokenPrivateKey: key}

	tests := map[string]string{
		"signed by another key": signIDToken(t, jose.RS256, newRSAKey(t), "jane@example.com"),
		"HMAC instead of RSA":   signIDToken(t, jose.HS256, []byte("0123456789abcdef0123456789abcdef"), "jane@example.com"),
		"not a token":           "garbage",
		"empty":                 "",
	}

	for name, token := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, http.StatusUnauthorized, callUserinfo(s, token).Code)
		})
	}
}

func TestOAuthJWKSVerifiesIssuedTokens(t *testing.T) {
	key := newRSAKey(t)
	s := &Service{OAuthIDTokenPrivateKey: key}
	rec := httptest.NewRecorder()

	s.oauthJWKS(rec, httptest.NewRequest(http.MethodGet, "/v1/oauth/jwks", nil))

	var jwks jose.JSONWebKeySet
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &jwks))
	require.Len(t, jwks.Keys, 1)

	token, err := jwt.ParseSigned(signIDToken(t, jose.RS256, key, "jane@example.com"), []jose.SignatureAlgorithm{jose.RS256})
	require.NoError(t, err)
	var claims idTokenClaims
	require.NoError(t, token.Claims(jwks.Keys[0].Key, &claims))
	assert.Equal(t, "jane@example.com", claims.Subject)
}
