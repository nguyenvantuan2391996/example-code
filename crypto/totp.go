package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func GenerateSecret() (string, error) {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(b), nil
}

func GenerateTOTP(secret string, period int64) (string, error) {
	return GenerateTOTPAt(secret, period, 0)
}

func GenerateTOTPAt(secret string, period int64, offset int) (string, error) {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", err
	}

	counter := (time.Now().Unix() / period) + int64(offset)

	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(counter))

	h := hmac.New(sha1.New, key)
	h.Write(buf)
	hash := h.Sum(nil)

	ot := hash[len(hash)-1] & 0x0F
	bin := (uint32(hash[ot])&0x7F)<<24 |
		(uint32(hash[ot+1])&0xFF)<<16 |
		(uint32(hash[ot+2])&0xFF)<<8 |
		(uint32(hash[ot+3]) & 0xFF)

	otp := bin % 1000000
	return fmt.Sprintf("%06d", otp), nil
}

func VerifyTOTP(secret string, code string, period int64, window int) bool {
	for i := -window; i <= window; i++ {
		g, _ := GenerateTOTPAt(secret, period, i)
		if g == code {
			return true
		}
	}
	return false
}

func TOTPURL(issuer, account, secret string) string {
	return fmt.Sprintf(
		"otpauth://totp/%s:%s?secret=%s&issuer=%s&period=30",
		issuer, account, secret, issuer,
	)
}

func Test() {
	secret, _ := GenerateSecret()
	fmt.Println("Secret:", secret)

	code, _ := GenerateTOTP(secret, 30)
	fmt.Println("TOTP:", code)

	valid := VerifyTOTP(secret, code, 30, 1)
	fmt.Println("Verify:", valid)

	url := TOTPURL("MyApp", "tuan@example.com", secret)
	fmt.Println("URL:", url)
}