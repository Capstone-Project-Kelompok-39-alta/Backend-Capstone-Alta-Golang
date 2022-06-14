package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
)

type invoceAuth struct {
	c    database.Config
	repo domains.AuthRepository
}

func NewInvoiceAuthService(repo domains.AuthRepository, c database.Config) *svcAuth {
	return &svcAuth{
		c:    c,
		repo: repo,
	}
}

func (s *svcAuth) invoceTelkom(numbertelkom int) error {
	user, _ := s.repo.LoginRepository(numbertelkom)

	if user.NumberTelkom != numbertelkom {
		return "your email error", http.StatusUnauthorized
	}

	return http.StatusOK
}

func (s *svcAuth) invocePLN(numberpln int) int {
	user, _ := s.repo.LoginRepository(numberpln)

	if user.Numberpln != numberpln {
		return "your email error", http.StatusUnauthorized
	}

	return http.StatusOK
}
