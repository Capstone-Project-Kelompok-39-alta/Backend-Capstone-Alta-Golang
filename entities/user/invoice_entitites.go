package user

type Invoice struct {
	NumberTelkom int `json:"telkomsel" validate:"required" gorm:"size:255"`
	NumberPLN    int `json:"pln" validate:"required" gorm:"size:255"`
}
