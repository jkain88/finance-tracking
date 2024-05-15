package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int    `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
