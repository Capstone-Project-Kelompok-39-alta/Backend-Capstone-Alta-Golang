package middleware

// Middleware user

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id int, secret string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["id"] = id
	claims["email"] = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
