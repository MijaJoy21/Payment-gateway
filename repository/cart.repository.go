package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateCart(ctx *gin.Context, data entity.Cart) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetCartByid(ctx *gin.Context, id int) ([]entity.Cart, error) {
	var data []entity.Cart

	query := db.DB.Preload("Product").Where("user_id = ?", id).Find(&data)

	// query := db.DB.Model(&data)
	// query = query.Where("id = ?", id)
	// query.Find(&data)

	return data, query.Error
}

func (db *repository) PutCart(ctx *gin.Context, id int, updatedData entity.Cart) error {
	var existingData entity.Cart

	query := db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates((updatedData))

	return query.Error
}
