package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (service *UserService) CreateUser(c *gin.Context) {
	var user models.User

	c.BindJSON(&user)

	c.JSON(http.StatusCreated, user)
	fmt.Println("Create User")
}
