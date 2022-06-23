package admin

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
)

type AuthRepository interface {
	LoginRepository(id_pegawai int) (admin entities.Admin, err error)
	RegisterRepository(admin entities.Admin) error
	GetUserRepository(name string) (entities.Admin, error)
}

type AuthService interface {
	LoginService(id_pegawai int, password string) (string, int)
	RegisterService(admin entities.Admin) error
	GetUserService(name string) (entities.Admin, error)
}
