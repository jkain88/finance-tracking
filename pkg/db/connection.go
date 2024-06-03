package db

import (
	"log"
	"time"

	"github.com/jkain88/finance-tracking/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://postgres:postgres@localhost:5432/finance-tracking?timezone=UTC"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Account{})
	db.AutoMigrate(&models.Transaction{})
	db.AutoMigrate(&models.Budget{})
	return db
}
