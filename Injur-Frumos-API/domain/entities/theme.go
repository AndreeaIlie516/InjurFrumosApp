package entities

import "gorm.io/gorm"

type Theme struct {
	gorm.Model
	Theme string `gorm:"column:theme;unique;not null" json:"theme" validate:"required"`
}
