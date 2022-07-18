package admin

import (
	"errors"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/testing/service_test/payment_status"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateInvoicePaymentStatusController(t *testing.T) {
	svc := payment_status.PaymentStatusService{}

	InvoicePaymentStatus := InvoicePaymentStatusController{
		Svc: &svc,
	}

	t.Run("StatusCreated", func(t *testing.T) {
		svc.On("CreateInvoicePaymentStatusService", mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/invoice-payment-status", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := InvoicePaymentStatus.CreateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		svc.On("CreateInvoicePaymentStatusService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/invoice-payment-status", io.Reader(strings.NewReader(`{"Status" : "Internal Server Error"}`)))
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := InvoicePaymentStatus.CreateInvoicePaymentStatusController(echoContext)
		if err != nil {
			return
		}
		assert.Equal(t, 500, w.Result().StatusCode)
	})
}

func TestGetAllInvoicePaymentStatusController(t *testing.T) {
	svc := payment_status.PaymentStatusService{}
	invoicePaymentStatus := new(payment_status.PaymentStatusService)
	invoicePaymentStatusData := []entities.InvoicePaymentStatus{
		{
			ID:   2,
			Name: "PENDING",
		},
		{
			ID:   3,
			Name: "PAID",
		},
		{
			ID:   4,
			Name: "EXPIRED",
		},
	}

	InvoicePaymentStatusController := InvoicePaymentStatusController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoicePaymentStatus.On("GetAllInvoicesPaymentStatusService").Return(invoicePaymentStatusData).Once()
		getAllStatusPayment := invoicePaymentStatus.GetAllInvoicesPaymentStatusService()
		svc.On("GetAllInvoicesPaymentStatusService").Return(invoicePaymentStatusData).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice-payment-status", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.GetAllInvoicePaymentStatusController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
		assert.Equal(t, getAllStatusPayment, invoicePaymentStatusData)
	})
}

func TestGetInvoicePaymentStatusByIDController(t *testing.T) {
	svc := payment_status.PaymentStatusService{}

	invoicePaymentStatus := new(payment_status.PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}
	InvoicePaymentStatusController := InvoicePaymentStatusController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, nil).Once()
		getInvoicePaymentStatusById, err := invoicePaymentStatus.GetInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		svc.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.GetInvoicePaymentStatusByIDController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, getInvoicePaymentStatusById, invoicePaymentStatusData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusNotFound", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, errors.New("error to make unit testing")).Once()
		getInvoicePaymentStatusById, err := invoicePaymentStatus.GetInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		svc.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice-payment-status/:id", io.Reader(strings.NewReader(`{"Status" : "Not Found"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.GetInvoicePaymentStatusByIDController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, getInvoicePaymentStatusById, invoicePaymentStatusData)
		assert.Equal(t, 404, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestUpdateInvoicePaymentStatusController(t *testing.T) {
	svc := payment_status.PaymentStatusService{}

	invoicePaymentStatus := new(payment_status.PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}
	InvoicePaymentStatusController := InvoicePaymentStatusController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(nil).Once()
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		svc.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/admin/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.UpdateInvoicePaymentStatusController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(errors.New("error to make unit testing")).Once()
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		svc.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("PUT", "/admin/invoice-payment-status/:id", io.Reader(strings.NewReader(`{"Status" : "Internal Server Error"}`)))
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.UpdateInvoicePaymentStatusController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestDeleteInvoicePaymentStatusController(t *testing.T) {
	svc := payment_status.PaymentStatusService{}

	invoicePaymentStatus := new(payment_status.PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}
	InvoicePaymentStatusController := InvoicePaymentStatusController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(nil).Once()
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		svc.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(nil).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/admin/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.DeleteInvoicePaymentStatusController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		svc.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/admin/invoice-payment-status/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoicePaymentStatusController.DeleteInvoicePaymentStatusController(echoContext)
		if er != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
