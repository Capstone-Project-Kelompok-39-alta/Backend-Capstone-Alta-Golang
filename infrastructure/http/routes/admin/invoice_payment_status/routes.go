package invoice_payment_status

import (
	invoicePaymentController "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	invoicePayment "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	invoicePaymentService "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)

	repo := invoicePayment.NewInvoicePaymentStatusRepository(db)
	svc := invoicePaymentService.NewInvoicePaymentStatusService(repo, conf)

	controller := invoicePaymentController.InvoicePaymentStatusController{
		Svc: svc,
	}

	InvoicePaymentStatusRoute := echo.Group("/admin/invoice-payment-status",
		middleware.CORS(),
		middleware.Logger(),
	)

	InvoicePaymentStatusRoute.POST("", controller.CreateInvoicePaymentStatusController, m.JWTTokenMiddleware())
	InvoicePaymentStatusRoute.GET("", controller.GetAllInvoicePaymentStatusController, m.JWTTokenMiddleware())
	InvoicePaymentStatusRoute.GET("/:id", controller.GetInvoicePaymentStatusByIDController, m.JWTTokenMiddleware())
	InvoicePaymentStatusRoute.PUT("/:id", controller.UpdateInvoicePaymentStatusController, m.JWTTokenMiddleware())
	InvoicePaymentStatusRoute.DELETE("/:id", controller.DeleteInvoicePaymentStatusController, m.JWTTokenMiddleware())
}
