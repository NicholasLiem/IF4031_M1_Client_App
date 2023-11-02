package utils

import (
	"errors"
	"strconv"
)

func VerifyUserId(UserID string) (uint, error) {
	userID, err := strconv.ParseUint(UserID, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse id")
	}
	return uint(userID), nil
}
