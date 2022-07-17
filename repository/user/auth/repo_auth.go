package auth

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domains.AuthRepositoryUser {
	return &repository{
		DB: db,
	}
}

func (r *repository) RegisterRepository(user entities.User) error {
	response := r.DB.Create(&user)

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (r *repository) LoginRepository(email string) (credential entities.User, err error) {
	err = r.DB.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&credential).Error

	return credential, nil
}
