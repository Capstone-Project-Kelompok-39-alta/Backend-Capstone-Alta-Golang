package user

import (
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
)

type AuthRepositoryUser interface {
	LoginRepository(email string) (user entities.User, err error)
	RegisterRepository(user entities.User) error
}

type AuthServiceUser interface {
	LoginService(email string, password string) (string, int)
	RegisterService(user entities.User) error
}
