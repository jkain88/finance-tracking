package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required" gorm:"unique"`
	Password  string `json:"password"`
	Provider  string `json:"provider"`
}
