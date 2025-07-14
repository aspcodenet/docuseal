package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model
	Name    string
	Content string
	UserID  uint
	Fields  []Field `gorm:"foreignKey:TemplateID"`
}

type Field struct {
	gorm.Model
	Name       string
	Type       string
	TemplateID uint
}
