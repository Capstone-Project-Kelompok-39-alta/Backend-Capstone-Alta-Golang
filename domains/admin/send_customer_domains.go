package admin

import "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"

type SendCustomerRepository interface {
	SendEmail(message entities.SendCustomer) error
}

type SendCustomerService interface {
	SendEmailService(message entities.SendCustomer) error
}
