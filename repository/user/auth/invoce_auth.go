package auth

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/user"
	entities "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities/user"
	"gorm.io/gorm"
)

type repositoryInvoice struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) domains.NewInvoiceRepository {
	return &repositoryInvoice{
		DB: db,
	}
}

func (r *repository) Telkom(NumberTelkom string) (credential entities.User, err error) {
	err = r.DB.Raw("SELECT * FROM users WHERE numbertelkom = ?", NumberTelkom).Scan(&credential).Error

	return credential, nil
}

func (r *repository) PLN(NumberPLN string) (credential entities.User, err error) {
	err = r.DB.Raw("SELECT * FROM users WHERE numberpln = ?", NumberPLN).Scan(&credential).Error

	return credential, nil
}
