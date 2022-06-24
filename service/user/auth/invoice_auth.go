package auth

import (
	_user "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/controller/user/auth"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
)

type invoceAuth struct {
	c    database.Config
	repo domains.AuthInvoiceRepository
}

func NewInvoiceAuthService(repo domains.AuthInvoiceRepository, c database.Config) *invoceAuth {
	return &invoceAuth{
		c:    c,
		repo: repo,
	}
}

func (s *invoceAuth) TelkomselInvoiceService(numbertelkom int) (*_user.Telkom, error) {
	user, err := s.repo.TelkomselInvoiceRepository(numbertelkom)

	if err != nil {
		return nil, err
	}

	userResponse := _user.InvoiceController(user)
	return &userResponse, nil
}
