package admin

import entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"

type InvoiceRepository interface {
	CreateInvoiceRepository(invoice entities.Invoice) error
	GetInvoiceUserRepository(issuerName string) (entities.Invoice, error)
	GetAllInvoiceRepository() (invoice []entities.Invoice, err error)
}

type InvoiceService interface {
	CreateInvoiceService(invoice entities.Invoice) error
	GetInvoiceUserService(issuerName string) (entities.Invoice, error)
	GetAllInvoiceService() (invoice []entities.Invoice, err error)
}
