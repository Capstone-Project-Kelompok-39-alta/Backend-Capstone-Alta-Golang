package entities

import "time"

type InvoiceItem struct {
	ID         int       `json:"id" form:"id"`
	ID_Invoice int       `json:"id_invoice" form:"id_invoice"`
	Product    string    `json:"product" form:"product"`
	Category   string    `json:"category" form:"category"`
	Qty        int       `json:"qty" form:"qty"`
	Price      float32   `json:"price" form:"price"`
	Subtotal   float32   `json:"subtotal" form:"subtotal"`
	Created_at time.Time `json:"created_at" form:"created_at"`
	Updated_at time.Time `json:"updated_at" form:"updated_at"`
}

func (*InvoiceItem) TableName() string {
	return "InvoiceItem"
}
