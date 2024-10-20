package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255;not null"`
	Email    string `json:"email" gorm:"size:255;unique;not null"`
	Password string `json:"-" gorm:"size:255;not null"`
	PhoneNo  string `json:"phone_no" gorm:"size:255;unique;not null"`
	Role     string `json:"role" gorm:"size:255;default:'user'"`
}
