import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSHA256(message, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
