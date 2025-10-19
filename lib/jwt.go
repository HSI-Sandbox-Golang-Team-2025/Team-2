package lib

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint `json:"id"`
}

func CreateJWT(userId uint) (string, error) {
	claims := Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return jwt.SignedString([]byte(GetEnv("JWT_SECRET", "")))
}

func ParseJwt(authHeader string) (*Claims, error) {
	var tokenString string

	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		return nil, errors.New("Invalid or expired token")
	}

	tokenString = authHeader[7:]

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			return []byte(GetEnv("JWT_SECRET", "")), nil
		},
	)

	if err != nil && !token.Valid {
		return nil, errors.New("Invalid or expired token")
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, errors.New("Invalid claims")
	}

	return claims, nil
}
