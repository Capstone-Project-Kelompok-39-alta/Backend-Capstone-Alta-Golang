package admin

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
)

type AuthRepository interface {
	LoginRepository(id_pegawai int) (admin entities.Admin, err error)
	RegisterRepository(admin entities.Admin) error
	GetUserRepository(id int) (entities.Admin, error)
	UpdateUserRepository(id int) (entities.Admin, error)
}

type AuthService interface {
	LoginService(id_pegawai int, password string) (string, int)
	RegisterService(admin entities.Admin) error
	GetUserService(id int) (entities.Admin, error)
	UpdateUserService(id int) (entities.Admin, error)
}
