package models

import (
	"gorm.io/gorm"
)

type AccountType string

const (
	Savings  AccountType = "Savings"
	Checking AccountType = "Checking"
	Credit   AccountType = "Credit"
)

func (t AccountType) IsValid() bool {
	switch t {
	case Savings, Checking, Credit:
		return true
	}
	return false
}

type Account struct {
	gorm.Model
	Name        string      `json:"name"`
	Type        AccountType `json:"type"`
	CreditLimit float64     `json:"credit_limit"`
	UserID      uint        `json:"user_id"`
	User        User
}
