package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) GetUsers(ctx *gin.Context) ([]entity.User, error) {
	var data []entity.User

	query := db.DB.Model(&data)
	query.Find(&data)

	return data, query.Error
}

func (db *repository) CreateUser(ctx *gin.Context, data entity.User) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetUserByEmail(ctx *gin.Context, email string) (entity.User, error) {
	var data entity.User
	query := db.DB.Model(&data)
	query = query.Where("email = ?", email)
	query.First(&data)

	return data, query.Error
}

func (db *repository) GetUserById(ctx *gin.Context, id int) (entity.User, error) {
	var data entity.User
	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}
