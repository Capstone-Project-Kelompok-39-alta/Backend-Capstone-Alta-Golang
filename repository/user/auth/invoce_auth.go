package auth

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"gorm.io/gorm"
)

type repositoryInvoice struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domains.AuthInvoiceRepository {
	return &repositoryInvoice{
		DB: db,
	}
}

func (r *repositoryInvoice) TelkomselInvoiceRepository(numbertelkom int) (user entities.InvoiceAdd, err error) {
	return
}

func (r *repositoryInvoice) PLNInvoiceRepository(numbertelkom int) (user entities.InvoiceAdd, err error) {
	return
}
