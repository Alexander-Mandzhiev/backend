package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func NewToken(username string, id int, secret string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":  id,
		"name": username,
		"exp":  time.Now().Add(duration).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
