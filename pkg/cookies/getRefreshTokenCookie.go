package cookies

import (
	"fmt"
	"net/http"
)

func GetRefreshTokenCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil || cookie == nil || cookie.Value == "" {
		return "", fmt.Errorf("missing or invalid refresh token")
	}
	return cookie.Value, nil
}
