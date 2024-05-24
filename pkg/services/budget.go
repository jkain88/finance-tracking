package services

import "gorm.io/gorm"

type BudgetService struct {
	db *gorm.DB
}

func NewBudgetService(db *gorm.DB) *BudgetService {
	return &BudgetService{
		db: db,
	}
}
