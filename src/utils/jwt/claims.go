package jwt

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func ParseUserIDClaim(ctx context.Context) (jwt.MapClaims, uint64, error) {
	claims, ok := ctx.Value("jwtClaims").(jwt.MapClaims)
	if !ok {
		return nil, 0, fmt.Errorf("JWT claims not found")
	}

	issuerUserIDStr, ok := claims["user_id"].(string)
	if !ok {
		return nil, 0, fmt.Errorf("JWT claim 'user_id' is not a string")
	}

	issuerUserID, err := strconv.ParseUint(issuerUserIDStr, 10, 64)
	if err != nil {
		return nil, 0, err
	}

	return claims, issuerUserID, nil
}
