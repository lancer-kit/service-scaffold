package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"strings"
)

// RandomString generates new random base32 Encoded random string.
func RandomString(length int, args ...string) (walletID string) {
	salt := make([]byte, 64)
	_, _ = rand.Read(salt)

	str := fmt.Sprintf("%s|%s", salt,
		strings.Join(args, "|"),
	)
	hash := sha256.Sum256([]byte(str))

	if length > 32 {
		length = 32
	}

	walletID = base32.StdEncoding.
		WithPadding(base32.NoPadding).
		EncodeToString(hash[:length])
	walletID = strings.ToLower(walletID)
	return
}
