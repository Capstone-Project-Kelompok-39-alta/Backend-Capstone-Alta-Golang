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

func (s *svcInvoice) CreateInvoiceService(Invoice entities.Invoice) error {
	return s.repo.CreateInvoiceRepository(Invoice)
}

func (s *svcInvoice) GetInvoiceService(Invoice entities.Invoice) {
}
