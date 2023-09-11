package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    jwtSecretKey    = os.Getenv("JWT_SECRET_KEY")
    AccessTokenTTL  = 15 * time.Minute
    RefreshTokenTTL = 24 * time.Hour
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
