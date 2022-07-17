package auth

import (
	"errors"
	domains2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestRegisterService(t *testing.T) {
	adminRegister := new(AuthService)
	adminRegisterData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	t.Run("Success", func(t *testing.T) {
		adminRegister.On("RegisterService", mock.Anything).Return(nil).Once()
		adminRegister := domains2.AuthService(adminRegister)
		err := adminRegister.RegisterService(adminRegisterData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		adminRegister.On("RegisterService", mock.Anything).Return(errors.New("error to make unit testing")).Once()

		adminRegister := domains2.AuthService(adminRegister)
		err := adminRegister.RegisterService(adminRegisterData)
		assert.Error(t, err)
	})
}

func TestLoginService(t *testing.T) {
	adminLogin := new(AuthService)
	adminLoginData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	t.Run("Success", func(t *testing.T) {
		adminLogin.On("LoginService", mock.Anything, mock.Anything).Return("success login", http.StatusOK).Once()
		adminLogin := domains2.AuthService(adminLogin)
		login, _ := adminLogin.LoginService(adminLoginData.IdPegawai, adminLoginData.Password)
		assert.Equal(t, login, "success login")
	})

	t.Run("Failed", func(t *testing.T) {
		adminLogin.On("LoginService", mock.Anything, mock.Anything).Return("Unauthorized", http.StatusUnauthorized).Once()
		adminLogin := domains2.AuthService(adminLogin)
		login, _ := adminLogin.LoginService(adminLoginData.IdPegawai, adminLoginData.Password)
		assert.Equal(t, login, "Unauthorized")
	})
}
