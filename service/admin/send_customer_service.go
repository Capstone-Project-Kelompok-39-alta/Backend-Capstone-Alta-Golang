package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/lib"
)

type svcSendCustomer struct {
	c    database.Config
	repo domains.SendCustomerRepository
}

func NewSendCustomerService(repo domains.SendCustomerRepository, c database.Config) domains.SendCustomerService {
	return &svcSendCustomer{
		c:    c,
		repo: repo,
	}
}

func (s *svcSendCustomer) SendEmailService(message entities.SendCustomer) error {
	err := lib.SendEmail(message)
	if err != nil {
		fmt.Println("this is error nya guys", err)
	}
	return s.repo.SendEmail(message)
}
