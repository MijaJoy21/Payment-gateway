package repository

import (
	"log"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (db *repository) CreateReview(ctx *gin.Context, data *entity.Review) error {
	query := db.DB.Model(&data)
	query = query.Save(&data)
	query.Find(&data)

	return query.Error
}

func (db *repository) GetReviewById(ctx *gin.Context, id int) (entity.Review, error) {
	var data entity.Review

	query := db.DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func (db *repository) PutReview(ctx *gin.Context, id int, updatedData entity.Review) error {
	var existingData entity.Review

	query := db.DB.Model(&existingData)
	query.Where("id = ?", id).Updates(updatedData)

	return query.Error
}

func (db *repository) DeleteReview(ctx *gin.Context, id int) error {
	var review entity.Review

	if err := db.DB.Where("id = ?", id).First(&review).Error; err != nil {
		log.Println("ID not Found")
		return err
	}
	query := db.DB.Delete(&review).Error

	return query
}
