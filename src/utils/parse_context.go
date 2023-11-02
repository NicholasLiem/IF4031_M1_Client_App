package utils

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/NicholasLiem/IF4031_M1_Client_App/internal/datastruct"
)

func ParseSessionUserFromContext(ctx context.Context) (*datastruct.SessionUserClient, error) {
	sessionData, ok := ctx.Value("jwtClaims").([]byte)
	if !ok {
		return nil, errors.New("session data not found in context")
	}

	var sessionUser datastruct.SessionUserClient
	if err := json.Unmarshal(sessionData, &sessionUser); err != nil {
		return nil, err
	}
	return &sessionUser, nil
}
