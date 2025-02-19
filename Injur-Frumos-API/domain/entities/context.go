package entities

import "gorm.io/gorm"

type Context struct {
	gorm.Model
	Context string `gorm:"column:context;unique;not null" json:"context" validate:"required"`
}
