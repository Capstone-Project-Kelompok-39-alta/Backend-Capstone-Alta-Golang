package upload_csv

import (
	uploadcsv2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	m "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	admin2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/repository/admin"
	admin3 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/service/admin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Routes(echo *echo.Echo, conf database.Config) {
	db := database.InitDB(conf)
	repo := admin2.NewUploadCsvRepository(db)
	repoInvoice := admin2.NewInvoiceRepository(db)
	svc := admin3.NewUploadCsvService(repo, conf)
	svcInvoice := admin3.NewInvoiceService(repoInvoice, conf)

	controllerUpload := uploadcsv2.UploadCsvController{
		Svc:        svc,
		SvcInvoice: svcInvoice,
	}

	controllerInvoice := uploadcsv2.InvoiceController{
		Svc: svcInvoice,
	}

	uploadcsvRoute := echo.Group("/admin",
		middleware.CORS(),
		middleware.Logger(),
	)

	uploadcsvRoute.POST("/upload_csv", controllerUpload.UploadCsvController, m.JWTTokenMiddleware())
	uploadcsvRoute.GET("/invoice", controllerInvoice.GetAllInvoice, m.JWTTokenMiddleware())
	uploadcsvRoute.GET("/invoice/:id", controllerInvoice.GetInvoiceUser, m.JWTTokenMiddleware())
}
