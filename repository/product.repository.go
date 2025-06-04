package repository

import (
	"log"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateProduct(ctx *gin.Context, data *entity.Product) error {
	query := db.DB.Model(&data)
	query = query.Save(&data)
	query.Find(&data)

	return query.Error
}

func (db *repository) GetProduct(ctx *gin.Context, params models.ParamsGetProduct) ([]entity.Product, int64, error) {
	var data []entity.Product
	var total int64

	query := db.DB.Model(&data)

	if params.Search != "" {
		query = query.Where("product.name like ?", "%"+params.Search+"%")
	}

	if params.CategoryId != 0 {
		query = query.Where("category_id = ?", params.CategoryId)
	}
	query = query.Joins("Category")
	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, total, query.Error

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
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}

func (db *repository) DeleteProduct(ctx *gin.Context, id int) error {
	var product entity.Product

	if err := db.DB.Where("id = ?", id).First(&product).Error; err != nil {
		log.Println("ID Not Found")
		return err
	}
	query := db.DB.Delete(&product).Error

	return query
}
