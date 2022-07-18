package send_customer

import (
	"errors"
	domains2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSendCustomerService_SendEmailService(t *testing.T) {
	sendEmail := new(SendCustomerService)
	sendEmailData := entities.SendCustomer{
		To:      "smanurung360@gmail.com",
		Subject: "Testing Subject",
		Body:    "Testing Body",
	}

	t.Run("Success", func(t *testing.T) {
		sendEmail.On("SendEmailService", mock.Anything).Return(nil).Once()
		sendEmail := domains2.SendCustomerService(sendEmail)
		err := sendEmail.SendEmailService(sendEmailData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		sendEmail.On("SendEmailService", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		sendEmail := domains2.SendCustomerService(sendEmail)
		err := sendEmail.SendEmailService(sendEmailData)

		assert.Error(t, err)
	})
}
