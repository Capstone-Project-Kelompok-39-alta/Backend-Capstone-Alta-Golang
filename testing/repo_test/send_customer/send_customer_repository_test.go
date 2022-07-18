package send_customer

import (
	"errors"
	domains2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSendEmail(t *testing.T) {
	sendEmail := new(SendCustomerRepository)
	sendEmailData := entities.SendCustomer{
		To:      "smanurung360@gmail.com",
		Subject: "Testing Subject",
		Body:    "Testing Body",
	}

	t.Run("Success", func(t *testing.T) {
		sendEmail.On("SendEmail", mock.Anything).Return(nil).Once()
		sendEmail := domains2.SendCustomerRepository(sendEmail)
		err := sendEmail.SendEmail(sendEmailData)

		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		sendEmail.On("SendEmail", mock.Anything).Return(errors.New("error to make unit testing")).Once()
		sendEmail := domains2.SendCustomerRepository(sendEmail)
		err := sendEmail.SendEmail(sendEmailData)

		assert.Error(t, err)
	})
}
