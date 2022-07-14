package entities

import "time"

type TransactionRecord struct {
	ID_Invoice         int       `json:"id_invoice"`
	ID_Invoice_Payment string    `json:"id_invoice_payment"`
	ID_User_Payment    string    `json:"id_user_payment"`
	Created_at         time.Time `json:"created_at"`
	Updated_at         time.Time `json:"updated_at"`
}

func (*TransactionRecord) TableName() string {
	return "TransactionRecord"
}
