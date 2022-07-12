package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
)

type repositoryAuth struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domains.AuthRepository {
	return &repositoryAuth{
		DB: db,
	}
}

func (r *repositoryAuth) RegisterRepository(admin entities.Admin) error {
	r.DB.Create(&admin)
	return nil
}

func (r *repositoryAuth) LoginRepository(id_Pegawai int) (credential entities.Admin, err error) {
	err = r.DB.Raw("SELECT * FROM admin WHERE id_pegawai = ?", id_Pegawai).Scan(&credential).Error

	return credential, nil
}

func (r *repositoryAuth) GetUserRepository(id int) (entities.Admin, error) {
	var admin entities.Admin
	r.DB.Where("id = ? ", id).Preload("admin").Find(&admin)
	return admin, nil
}

func (r *repositoryAuth) UpdateUserRepository(id int) (entities.Admin, error) {
	var admin entities.Admin
	r.DB.Where("id = ? ", id).Preload("admin").Find(&admin)
	return admin, nil
}
