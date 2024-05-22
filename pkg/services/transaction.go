package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

type TransactionInput struct {
	AccountID  uint    `json:"account_id" binding:"required"`
	CategoryID uint    `json:"category_id" binding:"required"`
	Name       string  `json:"name" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (service *TransactionService) CreateTransaction(c *gin.Context) {
	userId := c.GetUint("userId")

	var transactionInput TransactionInput
	err := c.ShouldBindJSON(&transactionInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	fmt.Println(userId)
}
