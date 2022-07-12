package entities

type Admin struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name" validate:"required" gorm:"size:32"`
	IdPegawai int    `json:"id_pegawai" validate:"required"`
	Email     string `json:"email" validate:"required" gorm:"size:255"`
	Password  string `json:"password" validate:"required" gorm:"size:255"`
}

type LoginAdmin struct {
	IdPegawai int    `json:"id_pegawai" validate:"required"`
	Password  string `json:"password" validate:"required" gorm:"column:password;size:255"`
}

type RegisterAdmin struct {
	Name      string `json:"name" validate:"required" gorm:"column:name;size:32"`
	IdPegawai int    `json:"id_pegawai" validate:"required"`
	Email     string `json:"email" validate:"required" gorm:"column:email;size:255"`
	Password  string `json:"password" validate:"required" gorm:"column:password;size:255"`
}

func (*Admin) TableName() string {
	return "admin"
}
