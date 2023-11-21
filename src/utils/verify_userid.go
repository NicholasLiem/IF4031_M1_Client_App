package utils

import (
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

func VerifyId(UserID string) (uint, *HttpError) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, &HttpError{
			Message:    "cannot parse id",
			StatusCode: http.StatusBadRequest,
		}
	}
	return uint(userID), nil
}

func VerifyUUID(UserID string) (uuid.UUID, *HttpError) {
	u, err := uuid.FromString(UserID)
	if err != nil {
		return uuid.UUID{}, &HttpError{
			Message:    "invalid UUID format",
			StatusCode: http.StatusBadRequest,
		}
	}
	return u, nil
}
