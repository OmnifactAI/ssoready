package saml_test

import (
	"math/big"
	"os"
	"testing"

	"github.com/ssoready/ssoready/internal/saml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Some identity providers issue certificates with negative serial numbers. Go
// rejects them by default since 1.23; go.mod keeps accepting them with
// `godebug x509negativeserial=1`. This test fails if that line is removed.
func TestParseMetadataAcceptsNegativeSerialCertificate(t *testing.T) {
	b, err := os.ReadFile("testdata/negative-serial-metadata.xml")
	require.NoError(t, err)

	res, err := saml.ParseMetadata(b)

	require.NoError(t, err)
	assert.Equal(t, "https://idp.example.com", res.IDPEntityID)
	assert.Equal(t, "https://idp.example.com/sso", res.RedirectURL)
	assert.Equal(t, big.NewInt(-12345), res.IDPCertificate.SerialNumber)
}
