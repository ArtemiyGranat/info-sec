package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
    jwtSecretKey    = []byte(os.Getenv("JWT_SECRET_KEY"))
    AccessTokenTTL  = 1 * time.Minute
    RefreshTokenTTL = 30 * 24 * time.Hour
)

func NewAccessToken(username string) string {
    claims := jwt.MapClaims{
        "sub": username, 
        "exp": time.Now().Add(AccessTokenTTL).Unix(),
    }
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    accessTokenString, _ := accessToken.SignedString(jwtSecretKey)

	return accessTokenString
}

func NewRefreshToken(username string) string {
    claims := jwt.MapClaims{
        "sub": username,
        "exp": time.Now().Add(RefreshTokenTTL).Unix(),
    }
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    refreshTokenString, _ := refreshToken.SignedString(jwtSecretKey)

	return refreshTokenString
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtSecretKey, nil
    })
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

    claims := token.Claims.(jwt.MapClaims)
    expirationTime, err := claims.GetExpirationTime()
    if err != nil {
        return nil, errors.New("Could not get an expiration time")
    }

    if expirationTime.Unix() < time.Now().Local().Unix() {
        return nil, errors.New("JWT is expired")
    }

	return token, nil

}