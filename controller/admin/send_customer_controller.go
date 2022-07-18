package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
	"net/http"
)

type SendCustomerController struct {
	Svc domains.SendCustomerService
}

// SendEmailController godoc
// @Summary Send Email to Customer
// @Description Admin can send email to customer for their invoice
// @Tags Send Customer
// @accept json
// @Produce json
// @Router /admin/send/email [post]
// @Param data body entities.SendCustomer true "required"
// @Success 201 {object} entities.SendCustomer
// @Failure 500 {object} entities.SendCustomer
// @Security JWT
func (co *SendCustomerController) SendEmailController(c echo.Context) error {
	messages := entities.SendCustomer{}
	err := c.Bind(&messages)

	sendEmail := co.Svc.SendEmailService(messages)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success to send email",
		"send":    sendEmail,
	})
}
