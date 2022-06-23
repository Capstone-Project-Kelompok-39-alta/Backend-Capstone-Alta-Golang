package user

type InvoiceAdd struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required" gorm:"size:32"`
	Telkomsel int    `json:"customeridtelkomsel" validate:"required" gorm:"size:255"`
	PLN       int    `json:"customeridpln" validate:"required" gorm:"size:255"`
}

type Telkom struct {
	Telkomsel int `json:"customeridtelkomsel" validate:"required"`
}

type Pln struct {
	PLN int `json:"customeridpln" validate:"required"`
}
