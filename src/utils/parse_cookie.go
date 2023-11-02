package utils

import (
	"errors"
	"github.com/gorilla/securecookie"
	"net/http"
	"os"
)

func ParseCookie(r *http.Request) (*string, error) {
	var (
		hashKey  = []byte(os.Getenv("HASH_KEY"))
		blockKey = []byte(os.Getenv("BLOCK_KEY"))
		s        = securecookie.New(hashKey, blockKey)
	)

	if sessionIdCookie, err := r.Cookie("sessionId"); err == nil {
		var sessionId string
		if err = s.Decode("sessionId", sessionIdCookie.Value, &sessionId); err == nil {
			return &sessionId, nil
		}
	}
	return nil, errors.New("fail to decode cookie")
}
