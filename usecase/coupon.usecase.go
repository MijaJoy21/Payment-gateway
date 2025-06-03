package usecase

import (
	"log"
	"math"
	"net/http"
	"payment-gateway/models"
	"payment-gateway/repository/entity"

	"github.com/gin-gonic/gin"
)

func (u *usecase) CreateCoupon(ctx *gin.Context, payload models.CreateCoupon) models.Response {
	res := models.Response{}

	data := entity.Coupon{
		Name:            payload.Name,
		Code:            payload.Code,
		MinimumPurchase: payload.MinimumPurchase,
		MaximumDiscount: payload.MaximumDiscount,
		Discount:        &payload.Discount,
		Type:            payload.Type,
		Quantity:        payload.Quantity,
	}

	if err := u.Repository.CreateCoupon(ctx, data); err != nil {
		log.Println("Error create coupon", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (u *usecase) GetListCoupon(ctx *gin.Context, params models.GetListCouponParams) models.Response {
	res := models.Response{}

	data, total, err := u.Repository.GetListCoupon(ctx, params)

	if err != nil {
		log.Println("Error get list coupon", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data
	res.Pagination = &models.Pagination{
		Page:      params.Page,
		Limit:     params.Limit,
		TotalData: total,
		LastPage:  int(math.Ceil(float64(total) / float64(params.Limit))),
	}
	return res
}
