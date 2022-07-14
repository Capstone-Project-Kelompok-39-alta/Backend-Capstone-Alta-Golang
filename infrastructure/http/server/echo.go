package server

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/constant"
	docs "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/docs"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	authAdmin "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/auth"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/invoice_payment_status"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/payment_gateway"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/send_email"
	uploadCsv "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/admin/upload_csv"
	authUser "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/routes/user/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"os"
	"time"
)

// @title Automatic Generate Invoice System API Documentation
// @description This is Automatic Generate Invoice API Documentation
// @version 2.0
// @host backend-capstone-alta-golang-staging.up.railway.app
// @BasePath
// @schemes https http
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
func Server() *echo.Echo {
	app := echo.New()
	conf := database.Config{}

	authAdmin.Routes(app, conf)
	uploadCsv.Routes(app, conf)
	send_email.Routes(app, conf)
	payment_gateway.Routes(app, conf)
	invoice_payment_status.Routes(app, conf)
	authUser.Routes(app, conf)

	app.Static(constant.STATIC_FILE_UPLOAD_CSV, constant.DIR_FILE_UPLOAD_CSV)
	app.GET("/swagger/*", echoSwagger.WrapHandler)

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
	app.Use(middleware.Recover())
	app.Use(middleware.Logger())
	app.Use(middleware.CORS())

	docs.SwaggerInfo.Host = os.Getenv("APP_HOST")

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Second,
	}
	if err := app.StartH2CServer(":8080", s); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return app
}
