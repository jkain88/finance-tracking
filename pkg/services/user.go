package services

import (
	"fmt"
	"net/http"

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

	c.JSON(http.StatusCreated, user)
	fmt.Println("Create User")
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
