package auth

// Service Login & Service

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"golang.org/x/crypto/bcrypt"
)

type svcAuth struct {
	c    database.Config
	repo domains.AuthRepository
}

func NewAuthService(repo domains.AuthRepository, c database.Config) *svcAuth {
	return &svcAuth{
		c:    c,
		repo: repo,
	}
}

func (s *svcAuth) RegisterService(user entities.User) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(password)
	return s.repo.RegisterRepository(user)
}

func (s *svcAuth) LoginService(email string, password string) (string, int) {
	user, _ := s.repo.LoginRepository(email)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if er != nil {
		return "wrong password", http.StatusUnauthorized
	}

	if user.Email != email {
		return "your email error", http.StatusUnauthorized
	}

	token, _ := middleware.CreateToken(int(user.ID), user.Email, s.c.JWT_KEY)
	return token, http.StatusOK
}
