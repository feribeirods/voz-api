package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"sub": userID,                                    // subject (identificador do usuário)
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(), // expira em 7 dias
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
