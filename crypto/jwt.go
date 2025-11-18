import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(secret string, claims jwt.MapClaims, expireIn time.Duration) (string, error) {
	claims["exp"] = time.Now().Add(expireIn).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyJWT(tokenStr, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
