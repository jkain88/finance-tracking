package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/db"
	"github.com/jkain88/finance-tracking/pkg/routes"
	"github.com/jkain88/finance-tracking/pkg/services"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func main() {
	db := db.Init()
	router := gin.Default()

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)

	// Initialize services with db connection
	userService := services.NewUserService(db)
	categoryService := services.NewCategoryService(db)
	accountService := services.NewAccountService(db)
	transactionService := services.NewTransactionService(db)
	budgetService := services.NewBudgetService(db)

	v1 := router.Group("/api/v1")
	{
		routes.UserRoutes(v1, userService)
		routes.CategoryRoutes(v1, categoryService)
		routes.AccountRoutes(v1, accountService)
		routes.TransactionRoutes(v1, transactionService)
		routes.BudgetRoutes(v1, budgetService)
		routes.AuthRoutes(v1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
