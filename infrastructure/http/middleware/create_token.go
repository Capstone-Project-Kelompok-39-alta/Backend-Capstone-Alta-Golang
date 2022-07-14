package middleware

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func CreateToken(id int, id_pegawai int, secret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["id"] = id
	claims["id_pegawai"] = id_pegawai

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func CreateTokenUser(id int, secret string, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["id"] = id
	claims["email"] = email
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
