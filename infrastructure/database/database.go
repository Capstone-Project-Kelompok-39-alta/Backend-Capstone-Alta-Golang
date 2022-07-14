package database

import (
	"fmt"
	entities2 "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	DB_HOST     string
	JWT_KEY     string
	BASE_URL    string
}

func InitDB(conf Config) *gorm.DB {
	conf = EnvDatabase()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	er := DB.AutoMigrate(
		&entities2.Admin{},
		&entities2.Invoice{},
		entities2.InvoicePaymentStatus{},
		entities2.InvoiceItem{},
		entities2.TransactionRecord{},
		entities2.SendCustomer{},
		entities2.User{},
	)

	if er != nil {
		panic(er)
	}

	return DB
}
