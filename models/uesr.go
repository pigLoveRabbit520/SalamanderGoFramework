package models

type User struct {
	Uid             int     `json:"id"`
	Phone           string  `gorm:"type:varchar(20);unique_index" json:"phone"`
	Password        string  `gorm:"type:varchar(32)" json:"-"`
}