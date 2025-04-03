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

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "b492243a9e39c8b36f4102c7b69c9c5bffedcd56f993f0da9a648b8d8925b2e7"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "4845646149dd7f3a97a8901ca25dcf8cd2dedcc8"
	)

	assert.Equal(t, privateKeyLen, len(privKey.Bytes()))

	address := privKey.Public().Address()
	assert.Equal(t, address.String(), addressStr)
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("some message")

	sig := privKey.Sign(msg)

	assert.True(t, sig.Verify(pubKey, msg))

	// test with invalid msg
	assert.False(t, sig.Verify(pubKey, []byte("wrong key")))

	invalidPrivateKey := GeneratePrivateKey()
	invalidPublicKey := invalidPrivateKey.Public()
	assert.False(t, sig.Verify(invalidPublicKey, msg))
}

func TestPublicKeyAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
}
