package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	response "github.com/NicholasLiem/IF4031_M1_Client_App/utils/http"
	"github.com/NicholasLiem/IF4031_M1_Client_App/utils/jwt"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionData")
		if err != nil {
			response.ErrorResponse(w, http.StatusBadRequest, "Cookie 'sessionData' is missing")
			return
		}

		tokenStr := cookie.Value
		claims, err := jwt.VerifyJWT(tokenStr)

		if err != nil {
			response.ErrorResponse(w, http.StatusUnauthorized, "Fail to verify JWT token: "+err.Error())
			return
		}

		claimsJSON, err := json.Marshal(claims)
		if err != nil {
			response.ErrorResponse(w, http.StatusInternalServerError, "Failed to marshal claims to JSON")
			return
		}

		ctx := context.WithValue(r.Context(), "jwtClaims", claimsJSON)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthenticateApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		ticketAuthHeader := "Bearer " + os.Getenv("TICKET_API_KEY")
		if authHeader != "" && (authHeader == ticketAuthHeader) {
			next.ServeHTTP(w, r)
		} else {
			response.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized access")
			return
		}
	})
}
