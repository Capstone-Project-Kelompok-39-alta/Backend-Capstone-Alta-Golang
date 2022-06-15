package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	Svc domains.AuthInvoiceService
}

func (co *InvoiceController) Invoice(c echo.Context) error {
	UserLogin := entities.Invoice{}

	err := c.Bind(&UserLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "bad request",
			"error":  err.Error(),
		})
	}

	statusCode := co.Svc.InvoiceService(UserLogin.NumberTelkom)

	if statusCode == http.StatusUnauthorized {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message":    "your email and password wrong",
			"statusCode": statusCode,
		})
	} else if statusCode == http.StatusInternalServerError {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message":    "internal server error",
			"statusCode": statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login user",
	})
}
