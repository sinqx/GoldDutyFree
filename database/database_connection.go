package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"sg/models"
	"time"
)

func Connect() *gorm.DB {
	username := ""        //os.Getenv("DB_USER")
	password := ""        // os.Getenv("DB_PASSWORD")
	dbName := "postgres"  // os.Getenv("DB_NAME")
	dbHost := "localhost" //os.Getenv("DB_HOST")
	dbType := "postgres"
	dbPort := "5432" // os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, dbHost, dbPort, dbName)

	database, err := gorm.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}
	database.Debug().AutoMigrate(models.User{})
	database.Debug().AutoMigrate(models.Order{})
	database.Debug().AutoMigrate(models.Gold{})

	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)
	database.DB().SetConnMaxLifetime(time.Hour)

	return database
}
