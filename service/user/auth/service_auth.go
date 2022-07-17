package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"golang.org/x/crypto/bcrypt"
)

type svcAuth struct {
	c    database.Config
	repo domains.AuthRepositoryUser
}

func NewAuthService(repo domains.AuthRepositoryUser, c database.Config) *svcAuth {
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

	token, _ := middleware.CreateTokenUser(int(user.ID), user.Email, s.c.JWT_KEY)
	return token, http.StatusOK
}
