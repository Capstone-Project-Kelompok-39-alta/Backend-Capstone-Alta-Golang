package user

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
)

type AuthInvoiceRepository interface {
	TelkomselInvoiceRepository(numbertelkom int) (user entities.InvoiceAdd, err error)
	PLNInvoiceRepository(numberpln int) (user entities.InvoiceAdd, err error)
}

type AuthInvoiceService interface {
	TelkomselInvoiceService(numbertelkom int) int
	PLNInvoiceService(numberpln int) int
}
