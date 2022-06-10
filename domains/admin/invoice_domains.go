package admin

import entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"

type InvoiceRepository interface {
	CreateInvoiceRepository(Invoice entities.Invoice) error
	GetInvoiceRepository(Invoice entities.Invoice)
}

type InvoiceService interface {
	CreateInvoiceService(Invoice entities.Invoice) error
	GetInvoiceService(Invoice entities.Invoice)
}
