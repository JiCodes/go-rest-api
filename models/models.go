package models

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Questions string `json:"question" gorm:"text; not null; default:null"`
	Answers string `json:"answer" gorm:"text; not null; default:null"`
}