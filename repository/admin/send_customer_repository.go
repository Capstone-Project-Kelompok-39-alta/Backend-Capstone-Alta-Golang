package admin

import (
	"fmt"
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
)

type sendCustomerRepository struct {
	DB *gorm.DB
}

func NewSendCustomerRepository(db *gorm.DB) domains.SendCustomerRepository {
	return &sendCustomerRepository{
		DB: db,
	}
}

func (r *sendCustomerRepository) SendEmail(message entities.SendCustomer) error {
	res := r.DB.Create(&message)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert to database")
	}
	return nil
}
