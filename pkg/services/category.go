package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/gorm"
)

type CategoryService struct {
	db *gorm.DB
}

type CategoryCreateInput struct {
	Name string `json:"name" binding:"required"`
}

func NewCategoryService(db *gorm.DB) *CategoryService {
	return &CategoryService{
		db: db,
	}
}

func (service *CategoryService) CreateCategory(c *gin.Context) {
	userId := c.GetUint("userId")
	var user models.User
	result := service.db.Find(&user, userId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	var categoryInput CategoryCreateInput
	err := c.ShouldBindJSON(&categoryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		User:   user,
		UserID: user.ID,
		Name:   categoryInput.Name,
	}

	result = service.db.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}
