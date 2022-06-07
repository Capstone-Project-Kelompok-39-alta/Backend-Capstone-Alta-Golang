package auth

import (
	"net/http"

	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
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

func (s *svcAuth) RegisterService(admin entities.User) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(password)
	return s.repo.RegisterRepository(admin)
}

func (s *svcAuth) LoginService(email string, password string) (string, int) {
	user, _ := s.repo.LoginRepository(email)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if er != nil {
		return "wrong password", http.StatusUnauthorized
	}

	// if (user.IdPegawai != email) || (user.Email < 8) {
	// 	return "your id_pegawai error", http.StatusUnauthorized
	// }

	// token, _ := middleware.CreateToken(int(user.ID), user.IdPegawai, s.c.JWT_KEY)
	return user.Email, http.StatusOK
	// return http.StatusOK
}
