package pkg

import (
	"os"
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

	return jwt.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
