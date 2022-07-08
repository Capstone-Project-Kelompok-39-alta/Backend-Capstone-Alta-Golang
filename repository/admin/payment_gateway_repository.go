package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
	"log"
)

type repositoryPaymentGateway struct {
	DB *gorm.DB
}

func NewPaymentGatewayRepository(db *gorm.DB) domains.PaymentGatewayRepository {
	return &repositoryPaymentGateway{
		DB: db,
	}
}

func (p *repositoryPaymentGateway) CreateTransactionRecord(id int, record entities.TransactionRecord) error {
	var transaction entities.TransactionRecord

	resp := p.DB.Where("id_invoice = ?", id).Find(&transaction)
	if resp.RowsAffected < 1 {
		resp1 := p.DB.Create(&record)
		if resp1.RowsAffected < 1 {
			log.Print("error insert")
			return fmt.Errorf("error to insert")
		}
	} else {
		resp2 := p.DB.Where("id_invoice = ?", id).UpdateColumns(&record)
		if resp2.RowsAffected < 1 {
			log.Print("error to update")
			return fmt.Errorf("error to update")
		}
	}

	return nil
}

func (p *repositoryPaymentGateway) GetIDInvoicePayment(id int) (record entities.TransactionRecord, err error) {
	resp := p.DB.Where("id_invoice = ?", id).Find(&record)
	if resp.RowsAffected < 1 {
		log.Printf("not found invoice")
		err = fmt.Errorf("not found the invoice")
	}
	return
}

func (p *repositoryPaymentGateway) GetInvoices(id int) (inv entities.Invoice, items []entities.InvoiceItem, err error) {
	resp := p.DB.Where("id = ?", id).Find(&inv)
	if resp.RowsAffected < 1 {
		log.Printf("not found the invoice")
		err = fmt.Errorf("not found the invoice")
	}

	resp1 := p.DB.Where("id_invoice = ?", id).Find(&items)
	if resp1.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}
	return
}

func (p *repositoryPaymentGateway) GetTotalAmount(id int) (float32, error) {
	var err error = nil
	var total float32 = 0
	var items []entities.InvoiceItem

	res := p.DB.Where("id_invoice = ?", id).Find(&items)
	if res.RowsAffected < 1 {
		log.Printf("not found invoice items")
		err = fmt.Errorf("not found invoice items")
	}

	for i := 0; i < len(items); i++ {
		total += items[i].Subtotal
	}

	return total, err
}

func (p *repositoryPaymentGateway) UpdateStatusInvoice(id int, invoice entities.Invoice) error {
	res := p.DB.Where("id = ?", id).UpdateColumns(&invoice)
	if res.RowsAffected < 1 {
		log.Print("Error Update")
		return fmt.Errorf("error update")
	}

	return nil
}
