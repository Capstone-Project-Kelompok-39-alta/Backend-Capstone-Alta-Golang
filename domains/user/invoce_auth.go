package user

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
)

type AuthInvoiceRepository interface {
	InvoiceRepository(numbertelkom int) (user entities.Invoice, err error)
}

type AuthInvoiceService interface {
	InvoiceService(numbertelkom int) int
}
