package models

import "gorm.io/gorm"

type TransactionType string

const (
	Add      TransactionType = "Add"
	Subtract TransactionType = "Subtract"
)

func (t TransactionType) IsValid() bool {
	switch t {
	case Add, Subtract:
		return true
	}
	return false
}

type Transaction struct {
	gorm.Model
	CategoryID uint
	Category   Category
	UserID     uint
	User       User
	AccountID  uint
	Account    Account
	Name       string
	Amount     float64
	Type       TransactionType
}
