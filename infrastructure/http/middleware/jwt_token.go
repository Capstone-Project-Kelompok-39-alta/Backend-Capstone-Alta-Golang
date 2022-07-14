package middleware

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTTokenMiddleware() echo.MiddlewareFunc {
	secret := database.Config{}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret.JWT_KEY),
		SigningMethod: "HS256",
	})
}

func JWTTokenMiddlewareUser() echo.MiddlewareFunc {
	secret := database.Config{}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(secret.JWT_KEY),
		SigningMethod: "HS256",
	})
}
