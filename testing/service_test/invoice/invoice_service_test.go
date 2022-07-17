package invoice

import (
	"errors"
	domains2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateInvoiceService(t *testing.T) {
	invoiceService := new(InvoiceService)
	invoiceData := entities.Invoice{
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

	t.Run("Success", func(t *testing.T) {
		invoiceService.On("CreateInvoiceService", mock.Anything).Return(nil).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		err := invoiceService.CreateInvoiceService(invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceService.On("CreateInvoiceService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		err := invoiceService.CreateInvoiceService(invoiceData)
		assert.Error(t, err)
	})
}

func TestGetAllInvoiceService(t *testing.T) {
	invoiceService := new(InvoiceService)
	invoiceData := []entities.Invoice{
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

	t.Run("Success", func(t *testing.T) {
		invoiceService.On("GetAllInvoiceService").Return(invoiceData, nil).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		getAllInvoice, err := invoiceService.GetAllInvoiceService()

		assert.Equal(t, getAllInvoice, invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceService.On("GetAllInvoiceService").Return(invoiceData, errors.New("error to make unit testing")).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		getAllInvoice, err := invoiceService.GetAllInvoiceService()

		assert.Equal(t, getAllInvoice, invoiceData)
		assert.Error(t, err)
	})
}

func TestGetInvoiceUserService(t *testing.T) {
	invoiceService := new(InvoiceService)
	invoiceData := entities.Invoice{
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

	t.Run("Success", func(t *testing.T) {
		invoiceService.On("GetInvoiceUserService", mock.Anything).Return(invoiceData, nil).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		getUserInvoice, err := invoiceService.GetInvoiceUserService(int(invoiceData.Id))

		assert.Equal(t, getUserInvoice, invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceService.On("GetInvoiceUserService", mock.Anything).Return(invoiceData, errors.New("error to make unit testing")).Once()
		invoiceService := domains2.InvoiceService(invoiceService)
		getUserInvoice, err := invoiceService.GetInvoiceUserService(int(invoiceData.Id))

		assert.Equal(t, getUserInvoice, invoiceData)
		assert.Error(t, err)
	})
}
