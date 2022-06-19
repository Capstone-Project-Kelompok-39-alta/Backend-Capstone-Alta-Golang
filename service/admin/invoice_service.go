package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
)

type svcInvoice struct {
	c    database.Config
	repo domains.InvoiceRepository
}

func NewInvoiceService(repo domains.InvoiceRepository, c database.Config) domains.InvoiceService {
	return &svcInvoice{
		c:    c,
		repo: repo,
	}
}

func (s *svcInvoice) CreateInvoiceService(invoice entities.Invoice) error {
	return s.repo.CreateInvoiceRepository(invoice)
}

func (s *svcInvoice) GetInvoiceUserService(id int) (entities.Invoice, error) {
	return s.repo.GetInvoiceUserRepository(id)
}

func (s *svcInvoice) GetAllInvoiceService() (invoice []entities.Invoice, err error) {
	return s.repo.GetAllInvoiceRepository()
}
