package auth_test

import (
	"errors"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/testing/repo_test/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRegisterRepository(t *testing.T) {
	adminRegister := new(auth.AuthRepository)
	adminRegsiterData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	t.Run("Success", func(t *testing.T) {
		adminRegister.On("RegisterRepository", mock.Anything).Return(nil).Once()

		adminRegister := domains.AuthRepository(adminRegister)
		err := adminRegister.RegisterRepository(adminRegsiterData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		adminRegister.On("RegisterRepository", mock.Anything).Return(errors.New("error to make unit testing")).Once()

		adminRegister := domains.AuthRepository(adminRegister)
		err := adminRegister.RegisterRepository(adminRegsiterData)
		assert.Error(t, err)
	})
}

func TestLoginRepository(t *testing.T) {
	adminLogin := new(auth.AuthRepository)
	adminLoginData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	t.Run("Success", func(t *testing.T) {
		adminLogin.On("LoginRepository", mock.Anything).Return(adminLoginData, nil).Once()
		adminLogin := domains.AuthRepository(adminLogin)
		login, err := adminLogin.LoginRepository(adminLoginData.IdPegawai)
		assert.Equal(t, login, adminLoginData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		adminLogin.On("LoginRepository", mock.Anything).Return(adminLoginData, errors.New("error to make unit testing")).Once()
		adminLogin := domains.AuthRepository(adminLogin)
		login, err := adminLogin.LoginRepository(adminLoginData.IdPegawai)
		assert.Equal(t, login, adminLoginData)
		assert.Error(t, err)
	})
}

func TestGetUserRepository(t *testing.T) {
	adminRepo := new(auth.AuthRepository)
	adminData := entities.Admin{
		Name:      "Name Testing",
		IdPegawai: 12345678,
		Email:     "emailTesting@gmail.com",
		Password:  "Password Testing",
	}

	t.Run("Success", func(t *testing.T) {
		adminRepo.On("GetUserRepository", mock.Anything).Return(adminData, nil).Once()
		adminRepo := domains.AuthRepository(adminRepo)
		adminUserRepo, err := adminRepo.GetUserRepository(int(adminData.IdPegawai))

		assert.Equal(t, adminUserRepo, adminData)
		assert.NoError(t, err)
	})

	t.Run("Failed", func(t *testing.T) {
		adminRepo.On("GetUserRepository", mock.Anything).Return(adminData, errors.New("error to get user by id_pegawai")).Once()
		adminRepo := domains.AuthRepository(adminRepo)
		adminUserRepo, err := adminRepo.GetUserRepository(int(adminData.IdPegawai))

		assert.Equal(t, adminUserRepo, adminData)
		assert.Error(t, err)
	})
}
