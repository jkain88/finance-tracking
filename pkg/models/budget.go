package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	UserID     uint
	User       User
	CategoryID uint
	Category   Category
	Label      string
	Amount     float64
}
