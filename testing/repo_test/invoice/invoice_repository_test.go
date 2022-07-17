package invoice

import (
	"errors"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateInvoiceRepository(t *testing.T) {
	invoiceRepo := new(InvoiceRepository)
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
		invoiceRepo.On("CreateInvoiceRepository", mock.Anything).Return(nil).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		err := invoiceRepo.CreateInvoiceRepository(invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceRepo.On("CreateInvoiceRepository", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		err := invoiceRepo.CreateInvoiceRepository(invoiceData)
		assert.Error(t, err)
	})
}

func TestGetAllInvoiceRepository(t *testing.T) {
	invoiceRepo := new(InvoiceRepository)
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
		invoiceRepo.On("GetAllInvoiceRepository").Return(invoiceData, nil).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		getAllInvoice, err := invoiceRepo.GetAllInvoiceRepository()
		assert.Equal(t, getAllInvoice, invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceRepo.On("GetAllInvoiceRepository").Return(invoiceData, errors.New("error to make unit testing")).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		getAllInvoice, err := invoiceRepo.GetAllInvoiceRepository()
		assert.Equal(t, getAllInvoice, invoiceData)
		assert.Error(t, err)
	})
}

func TestGetInvoiceUserRepository(t *testing.T) {
	invoiceRepo := new(InvoiceRepository)
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
		invoiceRepo.On("GetInvoiceUserRepository", mock.Anything).Return(invoiceData, nil).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		getInvoiceUser, err := invoiceRepo.GetInvoiceUserRepository(int(invoiceData.Id))

		assert.Equal(t, getInvoiceUser, invoiceData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoiceRepo.On("GetInvoiceUserRepository", mock.Anything).Return(invoiceData, errors.New("error to make unit testing")).Once()
		invoiceRepo := domains.InvoiceRepository(invoiceRepo)
		getInvoiceUser, err := invoiceRepo.GetInvoiceUserRepository(int(invoiceData.Id))

		assert.Equal(t, getInvoiceUser, invoiceData)
		assert.Error(t, err)
	})
}
