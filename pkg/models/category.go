package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name   string `json:"name" binding:"required"`
	UserID int
	User   User
}
