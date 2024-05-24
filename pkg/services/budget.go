package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/gorm"
)

type BudgetService struct {
	db *gorm.DB
}

type BudgetInput struct {
	CategoryID uint    `json:"category_id" binding:"required"`
	Label      string  `json:"label" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

func NewBudgetService(db *gorm.DB) *BudgetService {
	return &BudgetService{
		db: db,
	}
}

func (service *BudgetService) CreateBudget(c *gin.Context) {
	userId := c.GetUint("userId")

	var budgetInput BudgetInput
	err := c.ShouldBindJSON(&budgetInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget := models.Budget{
		UserID:     userId,
		CategoryID: budgetInput.CategoryID,
		Label:      budgetInput.Label,
		Amount:     budgetInput.Amount,
	}
	result := service.db.Create(&budget)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, budget)
}
