package repository

import (
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (d *repository) CreateCoupon(ctx *gin.Context, data entity.Coupon) error {
	query := d.DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func (d *repository) GetListCoupon(ctx *gin.Context, params models.GetListCouponParams) ([]entity.Coupon, int64, error) {
	var data []entity.Coupon
	var total int64
	query := d.DB.Model(&data)
	query.Count(&total)

	offset := (params.Page - 1) * params.Limit
	query = query.Limit(params.Limit).Offset(offset)
	query.Find(&data)

	return data, total, query.Error
}
