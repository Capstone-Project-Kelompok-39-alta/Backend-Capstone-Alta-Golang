package admin

import (
	"errors"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/testing/service_test/send_customer"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSendCustomerController_SendEmailController(t *testing.T) {
	svc := send_customer.SendCustomerService{}

	SendCustomerController := SendCustomerController{
		Svc: &svc,
	}

	t.Run("StatusCreated", func(t *testing.T) {
		svc.On("SendEmailService", mock.Anything).Return(nil).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/send/email", nil)
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := SendCustomerController.SendEmailController(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 201, w.Result().StatusCode)
	})

	t.Run("StatusInternalServerError", func(t *testing.T) {
		svc.On("SendEmailService", mock.Anything).Return(errors.New("error to make unit testing")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/admin/send/email", io.Reader(strings.NewReader(`{"Status" : "Internal Server Error"}`)))
		w := httptest.NewRecorder()

		echoContext := e.NewContext(r, w)
		err := SendCustomerController.SendEmailController(echoContext)

		if err != nil {
			return
		}

		assert.Equal(t, 500, w.Result().StatusCode)
	})
}
