package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PaymentGatewayController struct {
	Svc domains.PaymentGatewayService
}

// CreateXenditPaymentInvoiceController godoc
// @Summary Create Payment Invoice Using Xendit
// @Description Admin can create payment invoice using xendit
// @Tags InvoiceTransaction
// @accept json
// @Produce json
// @Router /admin/payment/xendit/invoice/{id} [post]
// @Param id path int true "id"
// @Success 201 {object} entities.TransactionRecord
// @Failure 500 {object} entities.TransactionRecord
// @Security JWT
func (co *PaymentGatewayController) CreateXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := co.Svc.CreateXenditPaymentInvoiceService(intID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "no id found",
			"error":   err,
		})
	}

	return c.JSONPretty(http.StatusCreated, resp, " ")
}

// GetXenditPaymentInvoiceController godoc
// @Summary Get Xendit Payment Invoice By ID
// @Description Admin can get xendit payment invoice by id
// @Tags InvoiceTransaction
// @accept json
// @Produce json
// @Router /admin/payment/xendit/invoice/{id} [get]
// @param id path int true "id"
// @Success 200 {object} entities.TransactionRecord
// @Failure 404 {object} entities.TransactionRecord
// @Security JWT
func (co *PaymentGatewayController) GetXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := co.Svc.GetXenditPaymentInvoiceService(intID)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err,
		})
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

// GetAllXenditPaymentInvoiceController godoc
// @Summary Get All Xendit Payment Invoice
// @Description Admin can get all xendit payment invoice
// @Tags InvoiceTransaction
// @accept json
// @Produce json
// @Router /admin/payment/xendit/invoice [get]
// @Success 200 {object} entities.TransactionRecord
// @Failure 404 {object} entities.TransactionRecord
// @Security JWT
func (co *PaymentGatewayController) GetAllXenditPaymentInvoiceController(c echo.Context) error {
	resp, err := co.Svc.GetAllXenditPaymentInvoiceService()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err,
		})
	}
	return c.JSONPretty(http.StatusOK, resp, " ")
}

// ExpireXenditPaymentInvoiceController godoc
// @Summary Expired Xendit Payment Invoice
// @Description Admin can expired xendit payment invoice
// @Tags InvoiceTransaction
// @accept json
// @Produce json
// @Router /admin/payment/xendit/invoice/expire/{id} [get]
// @Param id path int true "id"
// @Success 200 {object} entities.TransactionRecord
// @Failure 404 {object} entities.TransactionRecord
// @Security JWT
func (co *PaymentGatewayController) ExpireXenditPaymentInvoiceController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	resp, err := co.Svc.ExpireXenditPaymentInvoiceService(intID)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err,
		})
	}
	return c.JSONPretty(http.StatusOK, resp, " ")
}

// CallbackXenditPaymentInvoiceController godoc
// @Summary Xendit can callback by using this route
// @Description Xendit can callback by this route if customer is have been paying or the invoice is expired
// @Tags CallbackInvoice
// @accept json
// @Produce json
// @Router /admin/payment/xendit/invoice/callback [post]
// @Param data body entities.CallbackInvoice true "required"
// @Success 200 {object} entities.CallbackInvoice
// @Failure 404 {object} entities.CallbackInvoice
// @Security JWT
func (co *PaymentGatewayController) CallbackXenditPaymentInvoiceController(c echo.Context) error {
	invoiceCallback := entities.CallbackInvoice{}
	errors := c.Bind(&invoiceCallback)
	if errors != nil {
		return errors
	}
	fmt.Println("Invoice Callback : ", invoiceCallback)
	err := co.Svc.CallbackXenditPaymentInvoiceService(invoiceCallback)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "id not found",
			"error":   err,
		})
	}

	return c.JSONPretty(http.StatusOK, invoiceCallback, " ")
}
