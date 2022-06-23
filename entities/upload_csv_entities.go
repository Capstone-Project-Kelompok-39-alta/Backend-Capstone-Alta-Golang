package entities

type UploadCsv struct {
	File_csv string `json:"file_csv"`
}

func (*UploadCsv) TableName() string {
	return "upload_csv"
}
