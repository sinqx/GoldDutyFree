package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sg/models"
	"time"
)

func Connect() *gorm.DB {
	username := "DB_USERNAME"
	password := "POSTGRES_PASSWORD"
	dbName := "DB_DATABASE"
	dbHost := "DB_HOSTNAME"
	dbType := "DB_TYPE"
	dbPort := "DB_PORT"

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, username,
		password, dbName, dbPort) // connection string

	database, err := gorm.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}
	database.Debug().AutoMigrate(models.User{})
	database.Debug().AutoMigrate(models.Order{})
	database.Debug().AutoMigrate(models.OrderModel{})

	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
	database.DB().SetConnMaxLifetime(time.Hour)

	return database
}
