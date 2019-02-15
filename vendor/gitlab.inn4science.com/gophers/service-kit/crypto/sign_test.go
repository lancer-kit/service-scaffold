package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifySignature(t *testing.T) {
	privKey, pubKey := GenKeyPair()
	fmt.Println("Private Key: ", privKey)
	fmt.Println("Public  Key: ", pubKey)

	message := fmt.Sprintf("%s:%s", "4212340000", "test 42")
	fmt.Println(message)
	sig, err := SignMessage(privKey, message)
	assert.Equal(t, nil, err)

	fmt.Println("Signature: ", sig)

	ok, err := VerifySignature(pubKey, message, sig)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
}

func TestBase64EncodeNP(t *testing.T) {
	key := "this is key"
	out := Base64EncodeNP([]byte(key))
	println(out)
	in, err := Base64DecodeNP(out)
	assert.NoError(t, err)
	assert.Equal(t, string(in), key)
}

func TestBase32EncodeNP(t *testing.T) {
	key := "this is key"
	out := Base32EncodeNP([]byte(key))
	println(out)
	in, err := Base32DecodeNP(out)
	assert.NoError(t, err)
	assert.Equal(t, string(in), key)
}
