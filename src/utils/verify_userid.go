package utils

import (
	"net/http"
	"strconv"
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
