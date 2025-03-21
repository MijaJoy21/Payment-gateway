package repository

import (
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateProduct(ctx *gin.Context, data entity.Product) error {
	query := db.DB.Model(&data)
	query.Save(&data)

	return query.Error
}

func (db *repository) GetProduct(ctx *gin.Context) ([]entity.Product, error) {
	var data []entity.Product

	query := db.DB.Model(&data)
	query.Find(&data)

	return data, query.Error

}

func (db *repository) GetProductById(ctx *gin.Context, id int) (entity.Product, error) {
	var data entity.Product

	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (db *repository) PutProduct(ctx *gin.Context, id int, updatedData entity.Product) error {
	var existingData entity.Product

	query := db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates((updatedData))

	return query.Error
}
