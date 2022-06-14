package user

import (
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
)

type NewInvoiceRepository interface {
	InvoiceRepository(numbertelkom int) (user entities.Invoice, err error)
	RegisterRepository(user entities.User) error
}

type NewAuthService interface {
	InvoiceService(numbertelkom int) int
	RegisterService(user entities.User) error
}
