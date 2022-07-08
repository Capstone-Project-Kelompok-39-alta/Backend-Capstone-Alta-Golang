package admin

import (
	"fmt"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/constant"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
	"log"
	"strconv"
	"time"
)

type svcPaymentGateway struct {
	c    database.Config
	repo domains.PaymentGatewayRepository
}

func NewPaymentGatewayService(repo domains.PaymentGatewayRepository, c database.Config) domains.PaymentGatewayService {
	return &svcPaymentGateway{
		repo: repo,
		c:    c,
	}
}

func (s *svcPaymentGateway) CreateXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	inv, invItem, errRepo := s.repo.GetInvoices(id)

	if errRepo != nil {
		log.Print(errRepo)
		return nil, errRepo
	}

	if inv.ID_Payment_Status == 3 {
		return nil, fmt.Errorf("invoice already paid")
	}

	xendit.Opt.SecretKey = constant.XENDIT_SECRET_KEY

	customer := xendit.InvoiceCustomer{
		GivenNames:   inv.BuyerName,
		Email:        inv.BuyerEmail,
		MobileNumber: inv.BuyerPhone,
		Address:      inv.BuyerAddress,
	}

	var items []xendit.InvoiceItem

	for i := 0; i < len(invItem); i++ {
		item := xendit.InvoiceItem{
			Name:     invItem[i].Product,
			Quantity: invItem[i].Qty,
			Price:    float64(invItem[i].Price),
			Category: invItem[i].Category,
			Url:      "",
		}
		items = append(items, item)
	}

	var fees []xendit.InvoiceFee
	fee := xendit.InvoiceFee{
		Type:  "ADMIN",
		Value: 5000,
	}
	fees = append(fees, fee)

	NotificationType := []string{"email"}

	customerNotificationPreference := xendit.InvoiceCustomerNotificationPreference{
		InvoiceCreated:  NotificationType,
		InvoiceExpired:  NotificationType,
		InvoicePaid:     NotificationType,
		InvoiceReminder: NotificationType,
	}

	total, errors := s.repo.GetTotalAmount(inv.Id)

	if errors != nil {
		return nil, errors
	}

	for i := 0; i < len(fees); i++ {
		total += float32(fees[i].Value)
	}

	data := invoice.CreateParams{
		ExternalID:                     strconv.Itoa(inv.Id),
		Amount:                         float64(total),
		Description:                    inv.Description,
		InvoiceDuration:                86400,
		Customer:                       customer,
		CustomerNotificationPreference: customerNotificationPreference,
		SuccessRedirectURL:             "https://http.cat/200",
		FailureRedirectURL:             "https://http.cat/406",
		Currency:                       "IDR",
		Items:                          items,
		Fees:                           fees,
	}

	resp, err := invoice.Create(&data)
	if err != nil {
		log.Print(err)
		return resp, err
	}

	var statusInvoice entities2.Invoice
	statusInvoice.ID_Payment_Status = 2
	err2 := s.repo.UpdateStatusInvoice(inv.Id, statusInvoice)
	if err2 != nil {
		return nil, err2
	}

	transaction := entities2.TransactionRecord{
		ID_Invoice:         inv.Id,
		ID_Invoice_Payment: resp.ID,
		ID_User_Payment:    resp.UserID,
		Created_at:         time.Now(),
		Updated_at:         time.Now(),
	}

	errTransaction := s.repo.CreateTransactionRecord(inv.Id, transaction)
	if errTransaction != nil {
		log.Print(err)
		return resp, err
	}

	return resp, nil
}

func (s *svcPaymentGateway) GetXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	xendit.Opt.SecretKey = constant.XENDIT_SECRET_KEY

	ID, errRepo := s.repo.GetIDInvoicePayment(id)
	if errRepo != nil {
		log.Fatal(errRepo)
		return nil, errRepo
	}

	data := invoice.GetParams{
		ID: ID.ID_Invoice_Payment,
	}

	resp, err := invoice.Get(&data)

	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	return resp, nil
}

func (s *svcPaymentGateway) GetAllXenditPaymentInvoiceService() ([]xendit.Invoice, error) {
	xendit.Opt.SecretKey = constant.XENDIT_SECRET_KEY

	data := invoice.GetAllParams{
		Statuses: []string{"EXPIRED", "PENDING", "SETTLED"},
	}

	resps, err := invoice.GetAll(&data)
	if err != nil {
		log.Fatal(err)
		return resps, err
	}

	return resps, nil
}

func (s *svcPaymentGateway) ExpireXenditPaymentInvoiceService(id int) (*xendit.Invoice, error) {
	var statusInvoice entities2.Invoice

	xendit.Opt.SecretKey = constant.XENDIT_SECRET_KEY

	ID, errRepo := s.repo.GetIDInvoicePayment(id)
	if errRepo != nil {
		log.Fatal(errRepo)
		return nil, errRepo
	}

	data := invoice.ExpireParams{
		ID: ID.ID_Invoice_Payment,
	}

	resp, err := invoice.Expire(&data)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}

	statusInvoice.ID_Payment_Status = 4
	err2 := s.repo.UpdateStatusInvoice(id, statusInvoice)
	if err2 != nil {
		return nil, err2
	}

	return resp, nil
}

func (s *svcPaymentGateway) CallbackXenditPaymentInvoiceService(invoice entities2.CallbackInvoice) error {
	var statusInvoice entities2.Invoice

	StringID := invoice.ExternalID
	ID, _ := strconv.Atoi(StringID)
	if invoice.Status == "PENDING" {
		statusInvoice.ID_Payment_Status = 2
	} else if invoice.Status == "PAID" {
		statusInvoice.ID_Payment_Status = 3
	} else if invoice.Status == "EXPIRED" {
		statusInvoice.ID_Payment_Status = 4
	} else {
		return fmt.Errorf("status not found")
	}

	return s.repo.UpdateStatusInvoice(ID, statusInvoice)
}
