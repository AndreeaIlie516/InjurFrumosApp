package entities

import "gorm.io/gorm"

type Swear struct {
	gorm.Model
	RomanianSwear                 string `gorm:"column:romanian_swear;unique;not null" json:"romanian_swear" validate:"required"`
	EnglishWordByWordTranslations string `gorm:"column:english_word_by_word_translation;not null" json:"english_word_by_word_translation" validate:"required"`
	ActualMeaning                 string `gorm:"column:actual_meaning" json:"actual_meaning" validate:"required,max=256"`
	ContextualExample             string `gorm:"column:contextual_example" json:"contextual_example" validate:"required,max=256"`
	TypeID                        uint   `gorm:"column:type_id;not null" json:"type_id" validate:"required,number"`
	IntensityID                   uint   `gorm:"column:intensity_id;not null" json:"intensity_id" validate:"required,number"`
	ThemeID                       uint   `gorm:"column:theme_id;not null" json:"theme_id" validate:"required,number"`
	ContextID                     uint   `gorm:"column:context_id;not null" json:"context_id" validate:"required,number"`
	NoOfVotes                     int    `gorm:"column:no_of_votes" json:"no_of_votes" validate:"required,number"`
}
