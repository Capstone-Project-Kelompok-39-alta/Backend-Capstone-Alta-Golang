package entities

import "time"

type InvoiceItem struct {
	ID         int       `json:"id"`
	ID_Invoice int       `json:"id_invoice"`
	Product    string    `json:"product"`
	Category   string    `json:"category"`
	Qty        int       `json:"qty"`
	Price      float32   `json:"price"`
	Subtotal   float32   `json:"subtotal"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (*InvoiceItem) TableName() string {
	return "InvoiceItem"
}
