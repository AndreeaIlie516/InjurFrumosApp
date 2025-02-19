package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `gorm:"column:username;unique;not null" json:"username" binding:"required" validate:"required"`
	Password    string `gorm:"column:password;not null" json:"password" binding:"required" validate:"required"`
	FirstName   string `gorm:"column:first_name;not null" json:"first_name" binding:"required" validate:"required"`
	LastName    string `gorm:"column:last_name;not null" json:"last_name" binding:"required" validate:"required"`
	Email       string `gorm:"column:email;unique;not null" json:"email" binding:"required" validate:"required,email"`
	PhoneNumber string `gorm:"column:phone_number;unique;not null" json:"phone_number" binding:"required" validate:"required,e164"`
	Address     string `gorm:"column:address" json:"address" binding:"required" validate:"max=100"`
}
