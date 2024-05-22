package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/gorm"
)

type TransactionService struct {
	db *gorm.DB
}

type TransactionInput struct {
	AccountID  uint                   `json:"account_id" binding:"required"`
	CategoryID uint                   `json:"category_id" binding:"required"`
	Name       string                 `json:"name" binding:"required"`
	Amount     float64                `json:"amount" binding:"required"`
	Type       models.TransactionType `json:"type" binding:"required"`
}

func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

func (service *TransactionService) fetchRelatedData(userId, accountId, categoryId uint) (models.User, models.Account, models.Category, error) {
	var user models.User
	var account models.Account
	var category models.Category

	result := service.db.Find(&user, userId)
	if result.Error != nil {
		return user, account, category, errors.New(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return user, account, category, errors.New("user not found")
	}

	result = service.db.Where("user_id = ?", userId).Find(&account, accountId)
	if result.Error != nil {
		return user, account, category, errors.New(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return user, account, category, errors.New("account not found")
	}

	result = service.db.Where("user_id = ?", userId).Find(&category, categoryId)
	if result.Error != nil {
		return user, account, category, errors.New(result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return user, account, category, errors.New("category not found")
	}

	return user, account, category, nil
}

func (service *TransactionService) CreateTransaction(c *gin.Context) {
	userId := c.GetUint("userId")

	var transactionInput TransactionInput
	err := c.ShouldBindJSON(&transactionInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !transactionInput.Type.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type"})
		return
	}

	user, account, category, err := service.fetchRelatedData(userId, transactionInput.AccountID, transactionInput.CategoryID)
	fmt.Println("ERR", err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := models.Transaction{
		CategoryID: transactionInput.CategoryID,
		Category:   category,
		UserID:     userId,
		User:       user,
		AccountID:  transactionInput.CategoryID,
		Account:    account,
		Name:       transactionInput.Name,
		Type:       transactionInput.Type,
	}
	result := service.db.Create(&transaction)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	fmt.Println(userId, transactionInput)
	c.JSON(http.StatusCreated, transaction)
}
