package models

import "gorm.io/gorm"

type AccessToken struct {
	gorm.Model
	Token  string `gorm:"unique"`
	UserID uint
}
