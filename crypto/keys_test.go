package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, privateKeyLen, len(privKey.Bytes()))

	pubKey := privKey.Public()
	assert.Equal(t, publicKeyLen, len(pubKey.Bytes()))
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("some message")

	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(pubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("wrong key")))
}
