package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jkain88/finance-tracking/pkg/db"
	"github.com/jkain88/finance-tracking/pkg/models"
)

func PopulateTransactions() {
	//Define a list of transactions
	db := db.Init()
	transactions := []models.Transaction{
		{
			AccountID:  1,
			CategoryID: 1,
			UserID:     1,
			Name:       "Test Transaction 1",
			Type:       "debit",
			Amount:     100,
		},
		{
			AccountID:  2,
			CategoryID: 2,
			UserID:     1,
			Name:       "Test Transaction 2",
			Type:       "credit",
			Amount:     100,
		},
		{
			AccountID:  2,
			CategoryID: 2,
			UserID:     1,
			Name:       "Test Transaction 2",
			Type:       "credit",
			Amount:     100,
		},
		{
			AccountID:  2,
			CategoryID: 2,
			UserID:     1,
			Name:       "Test Transaction 2",
			Type:       "credit",
			Amount:     100,
		},
		{
			AccountID:  2,
			CategoryID: 2,
			UserID:     1,
			Name:       "Test Transaction 2",
			Type:       "credit",
			Amount:     100,
		},
		{
			AccountID:  2,
			CategoryID: 2,
			UserID:     1,
			Name:       "Test Transaction 2",
			Type:       "credit",
			Amount:     100,
		},
	}

	//Populate the transactions in the database
	for _, transaction := range transactions {
		fmt.Println(transaction)
		result := db.Create(&transaction)
		if result.Error != nil {
			panic(result.Error.Error())
		}
		transaction.CreatedAt = time.Now().AddDate(0, 0, rand.Intn(3))
		db.Save(&transaction)
	}
}
