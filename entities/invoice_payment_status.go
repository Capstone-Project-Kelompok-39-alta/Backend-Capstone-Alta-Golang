package entities

type InvoicePaymentStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (*InvoicePaymentStatus) TableName() string {
	return "InvoicePaymentStatus"
}
