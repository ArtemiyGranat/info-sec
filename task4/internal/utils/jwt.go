package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    jwtSecretKey    = os.Getenv("JWT_SECRET_KEY")
    AccessTokenTTL  = 24 * time.Hour
    RefreshTokenTTL = 30 * 24 * time.Hour
)

func NewAccessToken(username string) string {
    accessToken := jwt.New(jwt.SigningMethodHS256)
    accessToken.Claims = jwt.MapClaims{
        "sub": username, 
        "exp": time.Now().Add(AccessTokenTTL).Unix(),
    }
    accessTokenString, _ := accessToken.SignedString(jwtSecretKey)

	return accessTokenString
}

func NewRefreshToken(username string) string {
    refreshToken := jwt.New(jwt.SigningMethodHS256)
    refreshToken.Claims = jwt.MapClaims{
        "sub": username,
        "exp": time.Now().Add(RefreshTokenTTL).Unix(),
    }
    refreshTokenString, _ := refreshToken.SignedString(jwtSecretKey)

	return refreshTokenString
}

func ValidateRefreshToken(refreshTokenString string) (*jwt.Token, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecretKey, nil
    })
	if err != nil {
		return nil, err
	}

	if !refreshToken.Valid {
		return nil, errors.New("Invalid refresh token")
	}

	return refreshToken, nil

}