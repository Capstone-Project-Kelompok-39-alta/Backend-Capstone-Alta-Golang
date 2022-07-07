package send_email

import (
	admins "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	admin2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := admin.NewSendCustomerRepository(db)
	svc := admin2.NewSendCustomerService(repo, conf)

	controller := admins.SendCustomerController{
		Svc: svc,
	}

	sendEmailRoute := echo.Group("/admin",
		middleware.CORS(),
		middleware.Logger(),
	)

	sendEmailRoute.POST("/send/email", controller.SendEmailController, m.JWTTokenMiddleware())
}
