package jwt

import (
	"net/http"
	"strings"
)

func DecodeJWTToken(r *http.Request) string {
	tokenStr := r.Header.Get("Authorization")

	if len(tokenStr) == 0 {
		return ""
	}

	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)
	return tokenStr
}
