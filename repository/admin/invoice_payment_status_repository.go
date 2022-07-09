package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
)

type invoicePaymentStatusRepository struct {
	DB *gorm.DB
}

func NewInvoicePaymentStatusRepository(db *gorm.DB) domains.PaymentStatusRepository {
	return &invoicePaymentStatusRepository{
		DB: db,
	}
}

func (r *invoicePaymentStatusRepository) CreateInvoicePaymentStatus(payment entities.InvoicePaymentStatus) error {
	resp := r.DB.Create(&payment)
	if resp.RowsAffected < 1 {
		return fmt.Errorf("error to insert the data")
	}
	return nil
}

func (r *invoicePaymentStatusRepository) GetAllInvoicesPaymentStatus() []entities.InvoicePaymentStatus {
	invoicePaymentStatus := []entities.InvoicePaymentStatus{}
	r.DB.Find(&invoicePaymentStatus)

	return invoicePaymentStatus
}

func (r *invoicePaymentStatusRepository) GetInvoicePaymentStatusByID(id int) (payment entities.InvoicePaymentStatus, err error) {
	resp := r.DB.Where("id = ?", id).Find(&payment)
	if resp.RowsAffected < 1 {
		err = fmt.Errorf("not found the data")
	}

	return
}

func (r *invoicePaymentStatusRepository) UpdateInvoicePaymentStatusByID(id int, payment entities.InvoicePaymentStatus) error {
	resp := r.DB.Where("id = ?", id).UpdateColumns(&payment)
	if resp.RowsAffected < 1 {
		return fmt.Errorf("error to update the data")
	}

	return nil
}

func (r *invoicePaymentStatusRepository) DeleteInvoicePaymentStatusByID(id int) error {
	res := r.DB.Delete(&entities.InvoicePaymentStatus{ID: id})
	if res.RowsAffected < 1 {
		return fmt.Errorf("error to delete the data")
	}

	return nil
}
