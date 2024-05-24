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

func (service *BudgetService) UpdateBudget(c *gin.Context) {
	userId := c.GetUint("userId")
	id := c.Param("id")

	var budget models.Budget
	result := service.db.Where("user_id = ?", userId).Find(&budget, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}

	var budgetInput BudgetInput
	err := c.ShouldBindJSON(&budgetInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget.CategoryID = budgetInput.CategoryID
	budget.Label = budgetInput.Label
	budget.Amount = budgetInput.Amount
	result = service.db.Save(&budget)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, budget)
}

func (service *BudgetService) DeleteBudget(c *gin.Context) {
	userId := c.GetUint("userId")
	id := c.Param("id")

	var budget models.Budget
	result := service.db.Where("user_id = ?", userId).Find(&budget, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "budget not found"})
		return
	}

	result = service.db.Delete(&budget)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
