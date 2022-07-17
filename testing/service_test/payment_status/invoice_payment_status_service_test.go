package payment_status

import (
	"errors"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateInvoicePaymentStatusService(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("CreateInvoicePaymentStatusService", mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.CreateInvoicePaymentStatusService(invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("CreateInvoicePaymentStatusService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.CreateInvoicePaymentStatusService(invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestGetAllInvoicesPaymentStatusService(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusService)
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

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("GetAllInvoicesPaymentStatusService").Return(invoicePaymentStatusData).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		getAllStatusPayment := invoicePaymentStatus.GetAllInvoicesPaymentStatusService()
		assert.Equal(t, getAllStatusPayment, invoicePaymentStatusData)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("GetAllInvoicesPaymentStatusService").Return(invoicePaymentStatusData).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		getAllStatusPayment := invoicePaymentStatus.GetAllInvoicesPaymentStatusService()
		assert.Equal(t, getAllStatusPayment, invoicePaymentStatusData)
	})
}

func TestGetInvoicePaymentStatusByIDService(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, nil).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		getInvoicePaymentStatusByID, err := invoicePaymentStatus.GetInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		assert.Equal(t, getInvoicePaymentStatusByID, invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByIDService", mock.Anything).Return(invoicePaymentStatusData, errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		getInvoicePaymentStatusByID, err := invoicePaymentStatus.GetInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		assert.Equal(t, getInvoicePaymentStatusByID, invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestUpdateInvoicePaymentStatusByIDService(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByIDService", mock.Anything, mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestDeleteInvoicePaymentStatusByIDService(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusService)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByIDService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusService(invoicePaymentStatus)
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByIDService(int(invoicePaymentStatusData.ID))
		assert.Error(t, err)
	})
}
