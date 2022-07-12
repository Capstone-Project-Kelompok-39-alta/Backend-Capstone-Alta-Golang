package admin

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
)

type InvoiceRepository interface {
	CreateInvoiceRepository(invoice entities.Invoice) error
	GetInvoiceUserRepository(id int) (entities.Invoice, error)
	GetAllInvoiceRepository() (invoice []entities.Invoice, err error)
}

type InvoiceService interface {
	CreateInvoiceService(invoice entities.Invoice) error
	GetInvoiceUserService(id int) (entities.Invoice, error)
	GetAllInvoiceService() (invoice []entities.Invoice, err error)
}
