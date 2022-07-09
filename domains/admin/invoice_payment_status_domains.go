package admin

import "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"

type PaymentStatusRepository interface {
	CreateInvoicePaymentStatus(payment entities.InvoicePaymentStatus) error
	GetAllInvoicesPaymentStatus() []entities.InvoicePaymentStatus
	GetInvoicePaymentStatusByID(id int) (payment entities.InvoicePaymentStatus, err error)
	UpdateInvoicePaymentStatusByID(id int, payment entities.InvoicePaymentStatus) error
	DeleteInvoicePaymentStatusByID(id int) error
}

type PaymentStatusService interface {
	CreateInvoicePaymentStatusService(payment entities.InvoicePaymentStatus) error
	GetAllInvoicesPaymentStatusService() []entities.InvoicePaymentStatus
	GetInvoicePaymentStatusByIDService(id int) (payment entities.InvoicePaymentStatus, err error)
	UpdateInvoicePaymentStatusByIDService(id int, payment entities.InvoicePaymentStatus) error
	DeleteInvoicePaymentStatusByIDService(id int) error
}
