package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
)

func ParseSessionUserFromContext(ctx context.Context) (*datastruct.SessionUserClient, *HttpError) {
	sessionData, ok := ctx.Value("jwtClaims").([]byte)
	if !ok {
		return nil, &HttpError{
			Message:    "session data not found in context",
			StatusCode: http.StatusBadRequest,
		}
	}

	var sessionUser datastruct.SessionUserClient
	if err := json.Unmarshal(sessionData, &sessionUser); err != nil {
		return nil, &HttpError{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return &sessionUser, nil
}
