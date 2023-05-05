package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
	"sg/models"
	"time"
)

func Connect() *gorm.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbType := "postgres"
	dbPort := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, dbHost, dbPort, dbName)

	database, err := gorm.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}
	database.Debug().AutoMigrate(models.User{})
	database.Debug().AutoMigrate(models.Order{})
	database.Debug().AutoMigrate(models.OrderModel{})
	database.Debug().AutoMigrate(models.GoldModel{})

	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
	database.DB().SetConnMaxLifetime(time.Hour)

	return database
}
