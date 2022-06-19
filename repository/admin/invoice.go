package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"gorm.io/gorm"
)

type repositoryInvoice struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domains.InvoiceRepository {
	return &repositoryInvoice{
		DB: db,
	}
}

func (r *repositoryInvoice) CreateInvoiceRepository(invoice entities.Invoice) error {
	invoiceInsert := r.DB.Create(&invoice)

	if invoiceInsert.Error != nil {
		return invoiceInsert.Error
	}
	return nil
}

func (r *repositoryInvoice) GetInvoiceUserRepository(id int) (entities.Invoice, error) {
	var invoice entities.Invoice
	r.DB.Where("id = ?", id).Preload("invoice").Find(&invoice)
	return invoice, nil
}

func (r *repositoryInvoice) GetAllInvoiceRepository() (invoice []entities.Invoice, err error) {
	r.DB.Find(&invoice)

	return invoice, nil
}
