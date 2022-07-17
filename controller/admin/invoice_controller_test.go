package admin

import (
	"errors"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/testing/service_test/invoice"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetInvoiceUser(t *testing.T) {
	svc := invoice.InvoiceService{}
	invoiceService := new(invoice.InvoiceService)
	invoiceServiceData := entities.Invoice{
		Id:                1,
		IdCsvFile:         1,
		IssuerName:        "testing_user",
		IssuerEmail:       "testing_email",
		IssuerPhone:       "testing_phone",
		IssuerAddress:     "testing_issuer",
		BuyerName:         "testing_buyer",
		BuyerAddress:      "testing_buyer_address",
		BuyerPhone:        "testing_buyer_phone",
		BuyerEmail:        "testing_buyer_email",
		Tax:               5000,
		Total:             10000,
		ID_Payment_Status: 3,
		Description:       "testing_description",
		DueDate:           time.Now(),
	}

	InvoiceController := InvoiceController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoiceService.On("GetInvoiceUserService", mock.Anything).Return(invoiceServiceData, nil).Once()
		invoiceService, err := invoiceService.GetInvoiceUserService(int(invoiceServiceData.Id))
		svc.On("GetInvoiceUserService", mock.Anything).Return(invoiceServiceData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoiceController.GetInvoiceUser(echoContext)
		if er != nil {
			return
		}
		assert.Equal(t, invoiceService, invoiceServiceData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		invoiceService.On("GetInvoiceUserService", mock.Anything).Return(invoiceServiceData, errors.New("error to make unit testing")).Once()
		invoiceService, err := invoiceService.GetInvoiceUserService(int(invoiceServiceData.Id))
		svc.On("GetInvoiceUserService", mock.Anything).Return(invoiceServiceData, errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoiceController.GetInvoiceUser(echoContext)
		if er != nil {
			return
		}
		assert.Equal(t, invoiceService, invoiceServiceData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}

func TestGetAllInvoice(t *testing.T) {
	svc := invoice.InvoiceService{}
	invoiceService := new(invoice.InvoiceService)
	invoiceServiceData := []entities.Invoice{
		{
			Id:                2,
			IdCsvFile:         1,
			IssuerName:        "testing_user",
			IssuerEmail:       "testing_email",
			IssuerPhone:       "testing_phone",
			IssuerAddress:     "testing_issuer",
			BuyerName:         "testing_buyer",
			BuyerAddress:      "testing_buyer_address",
			BuyerPhone:        "testing_buyer_phone",
			BuyerEmail:        "testing_buyer_email",
			Tax:               5000,
			Total:             10000,
			ID_Payment_Status: 3,
			Description:       "testing_description",
			DueDate:           time.Now(),
		},
		{
			Id:                1,
			IdCsvFile:         2,
			IssuerName:        "testing_user",
			IssuerEmail:       "testing_email",
			IssuerPhone:       "testing_phone",
			IssuerAddress:     "testing_issuer",
			BuyerName:         "testing_buyer",
			BuyerAddress:      "testing_buyer_address",
			BuyerPhone:        "testing_buyer_phone",
			BuyerEmail:        "testing_buyer_email",
			Tax:               5000,
			Total:             10000,
			ID_Payment_Status: 2,
			Description:       "testing_description",
			DueDate:           time.Now(),
		},
	}

	InvoiceController := InvoiceController{
		Svc: &svc,
	}

	t.Run("StatusOK", func(t *testing.T) {
		invoiceService.On("GetAllInvoiceService").Return(invoiceServiceData, nil).Once()
		invoiceService, err := invoiceService.GetAllInvoiceService()
		svc.On("GetAllInvoiceService").Return(invoiceServiceData, nil).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoiceController.GetAllInvoice(echoContext)
		if er != nil {
			return
		}
		assert.Equal(t, invoiceService, invoiceServiceData)
		assert.Equal(t, 200, w.Result().StatusCode)
		assert.NoError(t, err)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		invoiceService.On("GetAllInvoiceService").Return(invoiceServiceData, errors.New("error to make unit testing")).Once()
		invoiceService, err := invoiceService.GetAllInvoiceService()
		svc.On("GetAllInvoiceService").Return(invoiceServiceData, errors.New("error to make unit testing")).Once()
		e := echo.New()
		r := httptest.NewRequest("GET", "/admin/invoice", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		er := InvoiceController.GetAllInvoice(echoContext)
		if er != nil {
			return
		}
		assert.Equal(t, invoiceService, invoiceServiceData)
		assert.Equal(t, 500, w.Result().StatusCode)
		assert.Error(t, err)
	})
}
