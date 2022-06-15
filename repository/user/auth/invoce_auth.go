package auth

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"gorm.io/gorm"
)

type repositoryInvoice struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domains.NewInvoiceRepository {
	return &repositoryInvoice{
		DB: db,
	}
}

func (r *repository) InvoiceRepository(numbertelkom int) (user entities.Invoice, err error) {
	return
}
