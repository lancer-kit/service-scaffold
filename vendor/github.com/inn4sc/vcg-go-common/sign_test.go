package vcgtools

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/inn4sc/vcg-go-common/currency"
)

func TestVerifySignature(t *testing.T) {
	privKey, pubKey := GenKeyPair()
	fmt.Println("Private Key: ", privKey)
	fmt.Println("Public  Key: ", pubKey)

	message := fmt.Sprintf("%s:%s", currency.Amount(4212340000).String(), "test 42")
	fmt.Println(message)
	sig, err := SignMessage(privKey, message)
	assert.Equal(t, nil, err)

	fmt.Println("Signature: ", sig)

	ok, err := VerifySignature(pubKey, message, sig)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, ok)
}
