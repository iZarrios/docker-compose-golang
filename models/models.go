package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model `json:"-"`
	Question string `gorm:"text;not null;default:null" json:"question"`
	Answer   string `gorm:"text;not null;default:null" json:"answer"`
}
