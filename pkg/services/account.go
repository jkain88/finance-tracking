package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/gorm"
)

type AccountService struct {
	db *gorm.DB
}

type AccountInput struct {
	Name        string             `json:"name" binding:"required"`
	Type        models.AccountType `json:"type" binding:"required"`
	CreditLimit float64            `json:"credit_limit"`
}

func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

func (service *AccountService) CreateAccount(c *gin.Context) {
	userId := c.GetUint("userId")
	var user models.User

	result := service.db.Find(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var accountInput AccountInput
	err := c.ShouldBindJSON(&accountInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !accountInput.Type.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	account := models.Account{
		UserID:      userId,
		User:        user,
		Name:        accountInput.Name,
		Type:        accountInput.Type,
		CreditLimit: accountInput.CreditLimit,
	}

	result = service.db.Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, account)
}

func (service *AccountService) UpdateAccount(c *gin.Context) {
	userId := c.GetUint("userId")
	accountId := c.Param("id")

	var account models.Account
	result := service.db.Where("id = ? AND user_id = ?", accountId, userId).Find(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var accountInput AccountInput
	err := c.ShouldBindJSON(&accountInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !accountInput.Type.IsValid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	account.Name = accountInput.Name
	account.Type = accountInput.Type
	result = service.db.Save(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, account)
}

func (service *AccountService) DeleteAccount(c *gin.Context) {
	userId := c.GetUint("userId")
	accountId := c.Param("id")

	var account models.Account
	result := service.db.Where("id = ? AND user_id = ?", accountId, userId).Find(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No account found with the given ID"})
		return
	}

	result = service.db.Delete(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
