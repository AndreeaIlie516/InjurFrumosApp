package entities

import "gorm.io/gorm"

type Type struct {
	gorm.Model
	Type string `gorm:"column:type;unique;not null" json:"type" validate:"required"`
}
