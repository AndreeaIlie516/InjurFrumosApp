package entities

import "gorm.io/gorm"

type Intensity struct {
	gorm.Model
	Intensity string `gorm:"column:intensity;unique;not null" json:"intensity" validate:"required"`
}
