package models

import "gorm.io/gorm"

type Submission struct {
	gorm.Model
	TemplateID uint
	UserID     uint
	Values     string // JSON string
}
