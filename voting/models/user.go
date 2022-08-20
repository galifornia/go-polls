package models

type User struct {
	Name     string `json:"name"`
	Password []byte `json:"-"`
	Email    string `json:"email" gorm:"unique"`
}
