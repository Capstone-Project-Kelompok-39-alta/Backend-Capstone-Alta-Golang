package admin

import (
	"errors"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
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
	idpegawai := admin.IdPegawai
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(password)

	if (admin.IdPegawai != idpegawai) || (admin.IdPegawai <= 8) {
		return errors.New("your id_pegawai error to created")
	}

	if admin.Name != "" {
		return errors.New("your name error to created")
	}

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

func (s *svcAuth) GetUserService(id_pegawai int) (entities.Admin, error) {
	return s.repo.GetUserRepository(id_pegawai)
}

func (s *svcAuth) UpdateUserService(id_pegawai int) (entities.Admin, error) {
	var admin entities.Admin
	passwords := admin.Password
	password, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(password)

	er := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(passwords))
	if er != nil {
		return entities.Admin{}, er
	}

	return s.repo.UpdateUserRepository(id_pegawai)
}
