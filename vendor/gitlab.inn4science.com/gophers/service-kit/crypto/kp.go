package crypto

import (
	"crypto/rand"

	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/crypto/ed25519"
)

// KP (KeyPair) is a representation of the `ed25519` key pair with common methods.
type KP struct {
	Private Key `json:"private" yaml:"private"`
	Public  Key `json:"public" yaml:"public"`
}

// RandomKP generates new KeyPair.
func RandomKP() KP {
	publicKey, privateKey, _ := ed25519.GenerateKey(rand.Reader)

	return KP{
		Private: Key(privateKey),
		Public:  Key(publicKey),
	}
}

// GenKeyPair generates new ed25519 private and public keys.
func GenKeyPair() (privateKey, publicKey string) {
	pubk, pk, _ := ed25519.GenerateKey(rand.Reader)

	privateKey = Base32Encode(pk)
	publicKey = Base32Encode(pubk)
	return
}

// SignData serializes the `data` and signs it using the `privateKey`.
func (kp *KP) SignData(data interface{}) (string, error) {
	message, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "can not to marshal data")
	}

	sig := ed25519.Sign(kp.Private.ToPrivate(), message)
	return Base32Encode(sig), nil
}

// Sign signs the `data`.
func (kp *KP) Sign(data []byte) (string, error) {
	sig := ed25519.Sign(kp.Private.ToPrivate(), data)
	return Base32Encode(sig), nil
}

// VerifySignature checks is valid `signature` of `message`.
func (kp *KP) VerifySignature(message, signature string) (bool, error) {
	rawSignature, err := Base32Decode(signature)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(kp.Public.ToPublicKey(), []byte(message), rawSignature), nil
}
