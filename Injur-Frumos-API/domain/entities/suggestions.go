package entities

import "gorm.io/gorm"

type Suggestions struct {
	gorm.Model
	SwearID uint   `gorm:"column:swear_id;not null" json:"swear_id" validate:"required,number"`
	UserID  uint   `gorm:"column:user_id;not null" json:"user_id" validate:"required,number"`
	Date    string `gorm:"column:date;not null" json:"date" validate:"required,date"`
	Status  int    `gorm:"column:status;not null" json:"status" validate:"required,number"`
}
