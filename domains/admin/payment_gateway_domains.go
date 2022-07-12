package admin

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/xendit/xendit-go"
)

type PaymentGatewayRepository interface {
	CreateTransactionRecord(id int, record entities.TransactionRecord) error
	GetIDInvoicePayment(id int) (record entities.TransactionRecord, err error)
	GetInvoices(id int) (inv entities.Invoice, items []entities.InvoiceItem, err error)
	GetTotalAmount(id int) (float32, error)
	UpdateStatusInvoice(id int, invoice entities.Invoice) error
}

type PaymentGatewayService interface {
	CreateXenditPaymentInvoiceService(id int) (*xendit.Invoice, error)
	GetXenditPaymentInvoiceService(id int) (*xendit.Invoice, error)
	GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error)
	ExpireXenditPaymentInvoiceService(int) (*xendit.Invoice, error)
	CallbackXenditPaymentInvoiceService(invoice entities.CallbackInvoice) error
}
