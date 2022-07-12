package payment_gateway

import (
	admin3 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	admin2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := admin.NewPaymentGatewayRepository(db)
	svc := admin2.NewPaymentGatewayService(repo, conf)

	controller := admin3.PaymentGatewayController{
		Svc: svc,
	}

	PaymentGatewayRoute := echo.Group("/admin/payment/xendit/invoice",
		middleware.Logger(),
		middleware.CORS(),
	)

	PaymentGatewayRoute.POST("/:id", controller.CreateXenditPaymentInvoiceController, m.JWTTokenMiddleware())
	PaymentGatewayRoute.GET("/:id", controller.GetXenditPaymentInvoiceController, m.JWTTokenMiddleware())
	PaymentGatewayRoute.GET("", controller.GetAllXenditPaymentInvoiceController, m.JWTTokenMiddleware())
	PaymentGatewayRoute.GET("/expire/:id", controller.ExpireXenditPaymentInvoiceController, m.JWTTokenMiddleware())
	PaymentGatewayRoute.POST("/callback", controller.CallbackXenditPaymentInvoiceController, m.JWTTokenMiddleware())
}
