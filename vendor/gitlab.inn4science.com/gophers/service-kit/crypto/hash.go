package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
)

// HashData serializes `data` and generate hash.
func HashData(data interface{}) (string, error) {
	str, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "can not to marshal data")
	}

	rawHash := sha256.Sum256(str)
	return base64.URLEncoding.EncodeToString(rawHash[:]), nil
}

// HashStrings joins passed strings and hash it.
func HashStrings(strs ...string) string {
	str := strings.Join(strs, "|")
	rawHash := sha256.Sum256([]byte(str))
	return base64.URLEncoding.EncodeToString(rawHash[:])
}
