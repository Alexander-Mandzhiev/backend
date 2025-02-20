package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateRefreshToken(token, secretKey string) (int, string, error) {
	claims, err := ParseToken(token, secretKey)
	if err != nil {
		return 0, "", fmt.Errorf("token is not valid: %w", err)
	}
	id, ok1 := claims["uid"].(float64)
	name, ok2 := claims["name"].(string)
	if !ok1 || !ok2 {
		return 0, "", fmt.Errorf("invalid data in token: uid or name is missing")
	}

	return int(id), name, nil
}

func ParseToken(tokenString string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signature method: %w", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token is not valid")
}
