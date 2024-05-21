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

type CategoryInput struct {
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

	var categoryInput CategoryInput
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

func (service *CategoryService) UpdateCategory(c *gin.Context) {
	userId := c.GetUint("userId")
	categoryId := c.Param("id")

	var category models.Category
	result := service.db.Find(&category, categoryId)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if category.UserID != userId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cannot update this category"})
		return
	}

	var categoryInput CategoryInput
	err := c.ShouldBindJSON(&categoryInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category.Name = categoryInput.Name
	service.db.Save(&category)
	c.JSON(http.StatusOK, category)
}
