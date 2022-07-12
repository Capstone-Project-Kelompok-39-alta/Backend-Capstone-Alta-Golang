package entities

import "time"

type Invoice struct {
	Id                int       `json:"id" csv:"id"`
	ID_Payment_Status int       `json:"id_payment_status" form:"id_payment_status" csv:"id_payment_status"`
	IdCsvFile         int       `json:"id_csv_file" csv:"id_csv"`
	Description       string    `json:"description" form:"description" csv:"description"`
	IssuerName        string    `json:"issuer_name" csv:"issuer_name"`
	IssuerEmail       string    `json:"issuer_email" csv:"issuer_email"`
	IssuerPhone       string    `json:"issuer_phone" csv:"issuer_phone"`
	IssuerAddress     string    `json:"issuer_address" csv:"issuer_address"`
	BuyerName         string    `json:"buyer_name" csv:"buyer_name"`
	BuyerPhone        string    `json:"buyer_phone" csv:"buyer_phone"`
	BuyerEmail        string    `json:"buyer_email" csv:"buyer_email"`
	BuyerAddress      string    `json:"buyer_address" csv:"buyer_address"`
	DueDate           time.Time `json:"due_date" csv:"due_date"`
	Tax               int       `json:"tax" csv:"tax"`
	Total             int       `json:"total" csv:"total"`
}

func (*Invoice) TableName() string {
	return "invoice"
}
