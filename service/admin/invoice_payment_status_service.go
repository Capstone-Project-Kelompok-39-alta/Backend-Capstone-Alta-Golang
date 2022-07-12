package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
)

type svcInvoicePaymentStatus struct {
	c    database.Config
	repo domains.PaymentStatusRepository
}

func NewInvoicePaymentStatusService(repo domains.PaymentStatusRepository, c database.Config) domains.PaymentStatusService {
	return &svcInvoicePaymentStatus{
		repo: repo,
		c:    c,
	}
}

func (s *svcInvoicePaymentStatus) CreateInvoicePaymentStatusService(payment entities.InvoicePaymentStatus) error {
	return s.repo.CreateInvoicePaymentStatus(payment)
}

func (s *svcInvoicePaymentStatus) GetAllInvoicesPaymentStatusService() []entities.InvoicePaymentStatus {
	return s.repo.GetAllInvoicesPaymentStatus()
}

func (s *svcInvoicePaymentStatus) GetInvoicePaymentStatusByIDService(id int) (payment entities.InvoicePaymentStatus, err error) {
	return s.repo.GetInvoicePaymentStatusByID(id)
}

func (s *svcInvoicePaymentStatus) UpdateInvoicePaymentStatusByIDService(id int, payment entities.InvoicePaymentStatus) error {
	return s.repo.UpdateInvoicePaymentStatusByID(id, payment)
}

func (s *svcInvoicePaymentStatus) DeleteInvoicePaymentStatusByIDService(id int) error {
	return s.repo.DeleteInvoicePaymentStatusByID(id)
}
