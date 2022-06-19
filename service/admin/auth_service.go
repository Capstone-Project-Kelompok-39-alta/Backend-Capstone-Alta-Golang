package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/http/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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

func (s *svcAuth) RegisterService(admin entities.Admin) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(password)

	return s.repo.RegisterRepository(admin)
}

func (s *svcAuth) LoginService(id_pegawai int, password string) (string, int) {
	user, _ := s.repo.LoginRepository(id_pegawai)

	er := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if er != nil {
		return "wrong password", http.StatusUnauthorized
	}

	if (user.IdPegawai != id_pegawai) || (user.IdPegawai < 8) {
		return "your id_pegawai error", http.StatusUnauthorized
	}

	token, _ := middleware.CreateToken(int(user.ID), user.IdPegawai, s.c.JWT_KEY)
	return token, http.StatusOK
}

func (s *svcAuth) GetUserService(name string) (entities.Admin, error) {
	return s.repo.GetUserRepository(name)
}
