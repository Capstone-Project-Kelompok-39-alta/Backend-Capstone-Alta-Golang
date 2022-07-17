package payment_status

import (
	"errors"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateInvoicePaymentStatus(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusRepository)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("CreateInvoicePaymentStatus", mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.CreateInvoicePaymentStatus(invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("CreateInvoicePaymentStatus", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.CreateInvoicePaymentStatus(invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestGetAllInvoiceStatus(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusRepository)
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
		invoicePaymentStatus.On("GetAllInvoicesPaymentStatus").Return(invoicePaymentStatusData).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		getAllInvoiceStatus := invoicePaymentStatus.GetAllInvoicesPaymentStatus()
		assert.Equal(t, getAllInvoiceStatus, invoicePaymentStatusData)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("GetAllInvoicesPaymentStatus").Return(invoicePaymentStatusData).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		getAllInvoiceStatus := invoicePaymentStatus.GetAllInvoicesPaymentStatus()
		assert.Equal(t, getAllInvoiceStatus, invoicePaymentStatusData)
	})
}

func TestGetInvoicePaymentStatusByID(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusRepository)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByID", mock.Anything).Return(invoicePaymentStatusData, nil).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		getInvoiceStatusById, err := invoicePaymentStatus.GetInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID))
		assert.Equal(t, getInvoiceStatusById, invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("GetInvoicePaymentStatusByID", mock.Anything).Return(invoicePaymentStatusData, errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		getInvoiceStatusById, err := invoicePaymentStatus.GetInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID))
		assert.Equal(t, getInvoiceStatusById, invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestUpdateInvoicePaymentStatusByID(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusRepository)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByID", mock.Anything, mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("UpdateInvoicePaymentStatusByID", mock.Anything, mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.UpdateInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID), invoicePaymentStatusData)
		assert.Error(t, err)
	})
}

func TestDeleteInvoicePaymentStatusByID(t *testing.T) {
	invoicePaymentStatus := new(PaymentStatusRepository)
	invoicePaymentStatusData := entities.InvoicePaymentStatus{
		ID:   1,
		Name: "Pending",
	}

	t.Run("Success", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByID", mock.Anything).Return(nil).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID))
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		invoicePaymentStatus.On("DeleteInvoicePaymentStatusByID", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		invoicePaymentStatus := domains.PaymentStatusRepository(invoicePaymentStatus)
		err := invoicePaymentStatus.DeleteInvoicePaymentStatusByID(int(invoicePaymentStatusData.ID))
		assert.Error(t, err)
	})
}
