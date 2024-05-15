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

	v1 := router.Group("/api/v1")
	{
		routes.UserRoutes(v1, userService)
	}
	fmt.Println("Hello")

	router.Run()
}
