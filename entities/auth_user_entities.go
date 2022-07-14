package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name" validate:"required" gorm:"size:32"`
	Email    string `json:"email" validate:"required" gorm:"size:255"`
	Password string `json:"password" validate:"required" gorm:"size:255"`
}

type LoginUser struct {
	Email    int    `json:"email" validate:"required"`
	Password string `json:"password" validate:"required" gorm:"column:password;size:255"`
}

type RegisterUser struct {
	Name     string `json:"name" validate:"required" gorm:"column:name;size:32"`
	Email    string `json:"email" validate:"required" gorm:"column:email;size:255"`
	Password string `json:"password" validate:"required" gorm:"column:password;size:255"`
}
