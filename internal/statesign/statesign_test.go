package statesign

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newSigner(b byte) *Signer {
	var key [32]byte
	for i := range key {
		key[i] = b
	}
	return &Signer{Key: key}
}

func TestEncodeDecodeRoundTrip(t *testing.T) {
	signer := newSigner(1)

	got, err := signer.Decode(signer.Encode(Data{SAMLFlowID: "saml_flow_123"}))

	require.NoError(t, err)
	assert.Equal(t, "saml_flow_123", got.SAMLFlowID)
}

func TestDecodeRejectsOtherKey(t *testing.T) {
	encoded := newSigner(1).Encode(Data{SAMLFlowID: "saml_flow_123"})

	_, err := newSigner(2).Decode(encoded)

	assert.ErrorContains(t, err, "digest mismatch")
}

func TestDecodeRejectsTamperedPayload(t *testing.T) {
	signer := newSigner(1)
	_, digest, _ := strings.Cut(signer.Encode(Data{SAMLFlowID: "saml_flow_123"}), ".")
	payload, _, _ := strings.Cut(signer.Encode(Data{SAMLFlowID: "saml_flow_456"}), ".")

	_, err := signer.Decode(payload + "." + digest)

	assert.ErrorContains(t, err, "digest mismatch")
}

func TestDecodeRejectsMalformedInput(t *testing.T) {
	signer := newSigner(1)

	for _, input := range []string{"", "no-dot", "!!!.AAAA", "AAAA.!!!"} {
		_, err := signer.Decode(input)
		assert.Error(t, err, "input %q", input)
	}
}
