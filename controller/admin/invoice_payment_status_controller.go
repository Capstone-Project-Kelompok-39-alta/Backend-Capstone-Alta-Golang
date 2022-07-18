package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type InvoicePaymentStatusController struct {
	Svc domains.PaymentStatusService
}

// CreateInvoicePaymentStatusController godoc
// @Summary Create Invoice Payment Status
// @Description Admin can create invoice payment for table invoice
// @Tags Invoice Payment Status
// @accept json
// @Produce json
// @Router /admin/invoice-payment-status [post]
// @param data body entities.InvoicePaymentStatus true "required"
// @Success 201 {object} entities.InvoicePaymentStatus
// @Failure 500 {object} entities.InvoicePaymentStatus
// @Security JWT
func (co *InvoicePaymentStatusController) CreateInvoicePaymentStatusController(c echo.Context) error {
	invoicePaymentStatus := entities.InvoicePaymentStatus{}
	err := c.Bind(&invoicePaymentStatus)

	invoicesPaymentStatus := co.Svc.CreateInvoicePaymentStatusService(invoicePaymentStatus)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":                "success",
		"invoice payment status": invoicesPaymentStatus,
	})
}

// GetAllInvoicePaymentStatusController godoc
// @Summary Get All Invoice Payment Status Information
// @Description Admin can get all invoice payment status information
// @Tags Invoice Payment Status
// @accept json
// @Produce json
// @Router /admin/invoice-payment-status [get]
// @Success 200 {object} entities.InvoicePaymentStatus
// @Security JWT
func (co *InvoicePaymentStatusController) GetAllInvoicePaymentStatusController(c echo.Context) error {
	invoicePaymentStatus := co.Svc.GetAllInvoicesPaymentStatusService()

	return c.JSONPretty(http.StatusOK, invoicePaymentStatus, " ")
}

// GetInvoicePaymentStatusByIDController godoc
// @Summary      Get Invoice Payment Status Information by Id
// @Description  Admin can get invoice payment status information by id
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /admin/invoice-payment-status/{id} [get]
// @param        id    path      int                         true  "id"
// @Success      200  {object}  entities.InvoicePaymentStatus
// @Failure      404  {object}  entities.InvoicePaymentStatus
// @Security     JWT
func (co *InvoicePaymentStatusController) GetInvoicePaymentStatusByIDController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoicePaymenStatus, err := co.Svc.GetInvoicePaymentStatusByIDService(intID)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found in database",
		})
	}

	return c.JSONPretty(http.StatusOK, invoicePaymenStatus, " ")
}

// UpdateInvoicePaymentStatusController godoc
// @Summary      Update Invoice Payment Status Information
// @Description  User can update invoice payment status information
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /admin/invoice-payment-status/{id} [put]
// @param        id   path      int  true  "id"
// @param        data  body      entities.InvoicePaymentStatus  true  "required"
// @Success      200  {object}  entities.InvoicePaymentStatus
// @Failure      500   {object}  entities.InvoicePaymentStatus
// @Security     JWT
func (co *InvoicePaymentStatusController) UpdateInvoicePaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	invoicePaymentStatus := entities.InvoicePaymentStatus{}
	err := c.Bind(&invoicePaymentStatus)

	updateInvoicesPayment := co.Svc.UpdateInvoicePaymentStatusByIDService(intID, invoicePaymentStatus)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "id can't change in database",
			"Data":    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":                "edited the data",
		"id":                     intID,
		"invoice payment status": updateInvoicesPayment,
	})
}

// DeleteInvoicePaymentStatusController godoc
// @Summary      Delete Invoice Payment Status Information
// @Description  Admin can delete invoice payment status information
// @Tags         Invoice Payment Status
// @accept       json
// @Produce      json
// @Router       /admin/invoice-payment-status/{id} [delete]
// @param        id   path      int  true  "id"
// @Success      200  {object}  entities.InvoicePaymentStatus
// @Failure      500  {object}  entities.InvoicePaymentStatus
// @Security     JWT
func (co *InvoicePaymentStatusController) DeleteInvoicePaymentStatusController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := co.Svc.DeleteInvoicePaymentStatusByIDService(intID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id to delete",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
		"id":      intID,
	})
}
