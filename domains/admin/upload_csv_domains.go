package admin

type UploadCsvRepository interface {
	UploadCsv(File_csv string)
}

type UploadCsvService interface {
	UploadCsv(File_csv string)
}
