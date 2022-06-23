package server

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/constant"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	authAdmin "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/auth"
	uploadCsv "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/upload_csv"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag/example/basic/docs"
	"os"
)

// @title Automatic Generate Invoice System API Documentation
// @description This is Automatic Generate Invoice API Documentation
// @version 2.0
// @host backend-capstone-alta-golang-staging.up.railway.app
// @BasePath
// @schemes http https
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	authAdmin.Routes(app, conf)
	uploadCsv.Routes(app, conf)
	app.Static(constant.STATIC_FILE_UPLOAD_CSV, constant.DIR_FILE_UPLOAD_CSV)
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")
	return app
}
