package auth

import (
	auths "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	auth2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := admin.NewAuthRepository(db)
	svc := auth2.NewAuthService(repo, conf)

	controller := auths.AuthController{
		Svc: svc,
	}

	authRoute := echo.Group("/admin",
		middleware.Logger(),
		middleware.CORS(),
	)

	authRoute.POST("/register", controller.Register)
	authRoute.POST("/login", controller.Login)
	authRoute.GET("/user/:id_pegawai", controller.GetUser, m.JWTTokenMiddleware())
	authRoute.PUT("/user/:id_pegawai", controller.UpdateUser, m.JWTTokenMiddleware())
}
