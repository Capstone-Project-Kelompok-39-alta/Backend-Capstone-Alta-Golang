package auth

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/admin"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domains.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) RegisterRepository(admin entities.Admin) error {
	response := r.DB.Create(&admin)

	if response.Error != nil {
		return response.Error
	}

	return nil
}

func (r *repository) LoginRepository(id_Pegawai int) (credential entities.Admin, err error) {
	err = r.DB.Raw("SELECT * FROM admin WHERE id_pegawai = ?", id_Pegawai).Scan(&credential).Error

	return credential, nil
}

func (r *repository) GetUserRepository(name string) (entities.Admin, error) {
	var admin entities.Admin
	r.DB.Where("name = ? ", name).First(&admin)
	return admin, nil
}
