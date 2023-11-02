package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	Token string `json:"token"`
}

func CreateJWT(userID, email, role string) (JWTToken, error) {

	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"email":     email,
		"role":      role,
		"logged_in": true,
		"exp":       time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)

	return JWTToken{Token: tokenString}, err // Fixed token structure and variable name
}

func VerifyJWT(tokenStr string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, nil // Return claims directly
}
