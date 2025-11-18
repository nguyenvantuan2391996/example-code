package crypto

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1(s string) string {
	sum := sha1.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}
