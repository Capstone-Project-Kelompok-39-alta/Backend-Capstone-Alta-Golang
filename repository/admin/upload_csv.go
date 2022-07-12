package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/entities"
	"gorm.io/gorm"
)

type uploadCsvRepository struct {
	DB *gorm.DB
}

func NewUploadCsvRepository(db *gorm.DB) domains.UploadCsvRepository {
	return &uploadCsvRepository{
		DB: db,
	}
}

func (r *uploadCsvRepository) UploadCsv(File_csv string) {
	uploadCsv := entities.UploadCsv{
		File_csv: File_csv,
	}
	r.DB.Create(&uploadCsv)
}
