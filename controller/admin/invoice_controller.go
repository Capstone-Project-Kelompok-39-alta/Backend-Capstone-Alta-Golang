package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type InvoiceController struct {
	Svc domains.InvoiceService
}

// GetInvoiceUser godoc
// @Summary Get Invoice User By ID
// @Description Admin can Get Invoice User By ID
// @Tags Invoice
// @accept json
// @Produce json
// @Router /admin/invoice/{id} [get]
// @param id path int true "id"
// @Success 200 {object} admin.Invoice
// @Failure 500 {object} admin.Invoice
// @Security JWT
func (co *InvoiceController) GetInvoiceUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	invoice, err := co.Svc.GetInvoiceUserService(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error to get the invoice",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get invoice by id",
		"data":    invoice,
	})
}

// GetAllInvoice godoc
// @Summary Get All Invoice
// @Description Admin Get All Invoice User
// @Tags Invoice
// @accept json
// @Produce json
// @Router /admin/invoice [get]
// @Success 200 {object} admin.Invoice
// @Failure 500 {object} admin.Invoice
// @Security JWT
func (co *InvoiceController) GetAllInvoice(c echo.Context) error {
	invoice, err := co.Svc.GetAllInvoiceService()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error to get all invoice",
			"status":  http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success to get all invoice",
		"data":    invoice,
	})
}
