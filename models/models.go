package models

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Questions string `json:"questions" gorm:"text; not null; default:null"`
	Answers string `json:"answers" gorm:"text; not null; default:null"`
}