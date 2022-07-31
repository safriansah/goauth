package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:VARCHAR;unique_index;NOT NULL"`
	Email    string `gorm:"type:VARCHAR;unique_index;NOT NULL"`
	Phone    string `gorm:"type:VARCHAR;unique_index;NOT NULL"`
	Password string `gorm:"type:TEXT;NOT NULL"`
}

func GetUser() User {
	var user User
	return user
}

func GetUsers() []User {
	var users []User
	return users
}
