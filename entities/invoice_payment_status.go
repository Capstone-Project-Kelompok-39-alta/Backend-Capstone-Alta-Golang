package entities

type InvoicePaymentStatus struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"id"`
}

func (*InvoicePaymentStatus) TableName() string {
	return "InvoicePaymentStatus"
}
