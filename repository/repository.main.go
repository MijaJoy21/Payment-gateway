package repository

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	repository struct {
		DB *gorm.DB
	}

	Repository interface {
	}
)

func InitRepository() Repository {
	return &repository{
		DB: InitDB(),
	}
}

func InitDB() *gorm.DB {

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	debug := os.Getenv("DB_DEBUG_MYSQL")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("Database Connect ERROR")
	}

	fmt.Println("mysql connected succesfully")
	if debug == "true" {
		return db.Debug()
	}

	return db
}
