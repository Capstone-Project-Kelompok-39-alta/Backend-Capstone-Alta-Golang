package admin

import entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"

type InvoiceRepository interface {
	GetInvoiceRepository(Invoice entities.Invoice)
}

type InvoiceService interface {
	GetInvoiceService(Invoice entities.Invoice)
}
