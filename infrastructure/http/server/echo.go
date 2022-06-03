package server

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	authAdmin "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/auth"
	"github.com/labstack/echo/v4"
)

func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	authAdmin.Routes(app, conf)

	return app
}
