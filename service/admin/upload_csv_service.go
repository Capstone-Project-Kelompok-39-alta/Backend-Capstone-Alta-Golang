package admin

import (
	domains "github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/domains/admin"
	"github.com/Capstone-Project-Kelompok-39-alta/Backend-Capstone-Alta-Golang/infrastructure/database"
)

type svcUploadCsv struct {
	c    database.Config
	repo domains.UploadCsvRepository
}

func NewUploadCsvService(repo domains.UploadCsvRepository, c database.Config) domains.UploadCsvService {
	return &svcUploadCsv{
		repo: repo,
		c:    c,
	}
}

func (s *svcUploadCsv) UploadCsv(File_csv string) {
	s.repo.UploadCsv(File_csv)
}
