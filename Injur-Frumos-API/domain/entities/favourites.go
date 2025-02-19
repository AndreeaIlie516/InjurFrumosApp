package entities

import "gorm.io/gorm"

type Favourites struct {
	gorm.Model
	SwearID uint `gorm:"column:swear_id;not null" json:"swear_id" validate:"required,number"`
	UserID  uint `gorm:"column:user_id;not null" json:"user_id" validate:"required,number"`
}
