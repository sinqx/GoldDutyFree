package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"sg/models"
	"time"
)

func Connect() *gorm.DB {
	username := "postgres"
	password := "663857"
	dbName := "postgres"
	dbHost := "db"
	dbType := "postgres"
	dbPort := "5432"

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, username,
		password, dbName, dbPort) // connection string

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
