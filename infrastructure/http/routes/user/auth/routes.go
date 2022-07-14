package auth

import (
	auths "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/user/auth"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/user/auth"
	auth2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/user/auth"
	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := auth.NewAuthRepository(db)
	svc := auth2.NewAuthService(repo, conf)

	controller := auths.ControllerUser{
		Svc: svc,
	}

	echo.POST("/user/register", controller.RegisterUser)
	echo.POST("/user/login", controller.LoginUser)
}
