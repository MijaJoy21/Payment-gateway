package repository

import (
	"fmt"
	"log"
	"os"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	repository struct {
		DB *gorm.DB
	}

	Repository interface {
		GetUsers(ctx *gin.Context) ([]entity.User, error)
		CreateUser(ctx *gin.Context, data entity.User) error
		GetUserByEmail(ctx *gin.Context, email string) (entity.User, error)
		GetUserById(ctx *gin.Context, id int) (entity.User, error)
		CreatePaymentMethod(ctx *gin.Context, data entity.PaymentMethod) error
		GetPaymentMethods(ctx *gin.Context) ([]entity.PaymentMethod, error)
		CreatePayment(ctx *gin.Context, data entity.Payment) error
		GetPaymentMethodById(ctx *gin.Context, id int) (entity.PaymentMethod, error)
		PutPaymentMethod(ctx *gin.Context, id int, updatedData entity.PaymentMethod) error
		CreateCategory(ctx *gin.Context, data entity.Category) error
		GetCategory(ctx *gin.Context) ([]entity.Category, error)
		GetCategoryById(ctx *gin.Context, id int) (entity.Category, error)
		PutCategory(ctx *gin.Context, id int, updatedData entity.Category) error
		CreateEtalase(ctx *gin.Context, data entity.Etalase) error
		GetEtalase(ctx *gin.Context) ([]entity.Etalase, error)
		GetEtalaseById(ctx *gin.Context, id int) (entity.Etalase, error)
		PutEtalase(ctx *gin.Context, id int, updatedData entity.Etalase) error
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
