package auth

import (
	auths "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin/auth"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin/auth"
	auth2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := auth.NewAuthRepository(db)
	svc := auth2.NewAuthService(repo, conf)

	controller := auths.AuthController{
		Svc: svc,
	}

	echo.POST("/admin/register", controller.Register)
	echo.POST("/admin/login", controller.Login)
}
