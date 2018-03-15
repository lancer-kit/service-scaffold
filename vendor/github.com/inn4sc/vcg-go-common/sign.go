package vcgtools

import (
	"crypto/rand"
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// GenKeyPair generates new ed25519 private and public keys.
func GenKeyPair() (privateKey, publicKey string) {
	pubk, pk, _ := ed25519.GenerateKey(rand.Reader)

	privateKey = base32Encode(pk)
	publicKey = base32Encode(pubk)
	return
}

// SignData serializes the `data` and signs it using the `privateKey`.
func SignData(privKey string, data interface{}) (string, error) {
	var err error
	pk := make(ed25519.PrivateKey, ed25519.PrivateKeySize)
	pk, err = base32Decode(privKey)
	if err != nil {
		return "", errors.Wrap(err, "invalid private key")
	}
	message, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "can not to marshal data")
	}

	msg := []byte(message)
	sig := ed25519.Sign(pk, msg)
	return base32Encode(sig), nil
}

// SignMessage signs the `message` using the `privateKey`.
func SignMessage(privKey, message string) (string, error) {
	var err error
	pk := make(ed25519.PrivateKey, ed25519.PrivateKeySize)
	pk, err = base32Decode(privKey)
	if err != nil {
		return "", err
	}

	msg := []byte(message)
	sig := ed25519.Sign(pk, msg)
	return base32Encode(sig), nil
}

// VerifySignature checks is valid `signature` of `message`.
func VerifySignature(pubKey, message, signature string) (bool, error) {
	var err error
	var rawSignature []byte
	var rawPubKey = make(ed25519.PublicKey, ed25519.PublicKeySize)

	rawPubKey, err = base32Decode(pubKey)
	if err != nil {
		return false, err
	}
	rawSignature, err = base32Decode(signature)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(rawPubKey, []byte(message), rawSignature), nil
}
