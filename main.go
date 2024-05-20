package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/db"
	"github.com/jkain88/finance-tracking/pkg/routes"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func main() {
	db := db.Init()
	router := gin.Default()

	// Initialize services with db connection
	userService := services.NewUserService(db)
	categoryService := services.NewCategoryService(db)

	v1 := router.Group("/api/v1")
	{
		routes.UserRoutes(v1, userService)
		routes.CategoryRoutes(v1, categoryService)
	}
	fmt.Println("Hello")

	router.Run()
}
