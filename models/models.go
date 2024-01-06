package models

import (
	"gorm.io/gorm"
)

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text; not null; default:null"`
	Answer string `json:"answer" gorm:"text; not null; default:null"`
}

// The Fact struct embeds gorm.Model, which itself includes the fields for the standard GORM model, including ID, CreatedAt, UpdatedAt, and DeletedAt. By embedding gorm.Model, your Fact struct inherits these fields:
// ID: The primary key of the record.
// CreatedAt: The timestamp when the record was created.
// UpdatedAt: The timestamp when the record was last updated.
// DeletedAt: The timestamp when the record was soft deleted (if soft delete is enabled).