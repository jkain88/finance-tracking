package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"github.com/jkain88/finance-tracking/pkg/utils"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserProfile struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
}

type UserCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UserAccount struct {
	ID   uint               `json:"id"`
	Name string             `json:"name"`
	Type models.AccountType `json:"type"`
}

type UserTransaction struct {
	ID       uint                   `json:"id"`
	Name     string                 `json:"name"`
	Type     models.TransactionType `json:"type"`
	Account  UserAccount            `json:"account"`
	Category UserCategory           `json:"category"`
}

type UserBudget struct {
	ID       uint         `json:"id"`
	Category UserCategory `json:"category"`
	Label    string       `json:"label"`
	Amount   float64      `json:"amount"`
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (service *UserService) CreateUser(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword
	result := service.db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	// Create user initial categories
	for _, categoryName := range utils.Categories {
		category := models.Category{
			Name:   categoryName,
			UserID: user.ID,
		}
		result := service.db.Create(&category)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
			return
		}
	}

	c.JSON(http.StatusCreated, user)
}

func (service *UserService) SignIn(c *gin.Context) {
	var input SignInInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := service.db.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	isPasswordVerified := utils.CheckPassword(input.Password, user.Password)
	if !isPasswordVerified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email/password."})
		return
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (service *UserService) Me(c *gin.Context) {
	userId := c.GetUint("userId")

	var user models.User
	result := service.db.Find(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	userProfile := UserProfile{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	c.JSON(http.StatusOK, userProfile)
}

func (service *UserService) UserCategories(c *gin.Context) {
	var categories []models.Category
	var userCategories []UserCategory
	userId := c.GetUint("userId")

	result := service.db.Where("user_id = ?", userId).Find(&categories)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for _, category := range categories {
		userCategories = append(userCategories, UserCategory{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	c.JSON(http.StatusOK, userCategories)
}

func (service *UserService) UserAccounts(c *gin.Context) {
	var accounts []models.Account
	var userAccounts []UserAccount
	userId := c.GetUint("userId")

	result := service.db.Where("user_id = ?", userId).Find(&accounts)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for _, account := range accounts {
		userAccounts = append(userAccounts, UserAccount{
			ID:   account.ID,
			Name: account.Name,
			Type: account.Type,
		})
	}

	c.JSON(http.StatusOK, userAccounts)
}

func (service *UserService) UserTransactions(c *gin.Context) {
	var transactions []models.Transaction
	var userTransactions []UserTransaction
	var pagination utils.PaginationResult
	userId := c.GetUint("userId")

	db := service.db
	db = db.Scopes(utils.Paginate(c, db, transactions, &pagination))
	db = db.Order("id desc")
	db = db.Where("transactions.user_id = ?", userId)
	db = db.Preload("Account")

	account := c.Query("account")
	if account != "" {
		db = db.Joins("JOIN accounts ON accounts.id = transactions.account_id").Where("accounts.name = ?", account)
	}

	category := c.Query("category")
	if category != "" {
		db = db.Joins("Category", service.db.Where(&models.Category{Name: category}))
	} else {
		db = db.Joins("Category")
	}

	result := db.Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for _, transaction := range transactions {
		fmt.Println(transaction)
		userTransactions = append(userTransactions, UserTransaction{
			ID:   transaction.ID,
			Name: transaction.Name,
			Category: UserCategory{
				ID:   transaction.CategoryID,
				Name: transaction.Category.Name,
			},
			Account: UserAccount{
				ID:   transaction.AccountID,
				Name: transaction.Account.Name,
				Type: transaction.Account.Type,
			},
			Type: transaction.Type,
		})
	}

	pagination.Rows = userTransactions

	c.JSON(http.StatusOK, pagination)
}

func (service *UserService) UserBudgets(c *gin.Context) {
	var budgets []models.Budget
	var userBudgets []UserBudget
	userId := c.GetUint("userId")

	result := service.db.Preload("Category").Where("user_id = ?", userId).Find(&budgets)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	for _, budget := range budgets {
		userBudgets = append(userBudgets, UserBudget{
			ID: budget.ID,
			Category: UserCategory{
				ID:   budget.CategoryID,
				Name: budget.Category.Name,
			},
			Label:  budget.Label,
			Amount: budget.Amount,
		})
	}

	c.JSON(http.StatusOK, userBudgets)
}
