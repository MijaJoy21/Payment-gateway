package controllers

import (
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateCoupon(ctx *gin.Context) {
	log.Println("<<<<CreateCouponControllers>>>")
	var res models.Response
	payload := models.CreateCoupon{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error bad request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helpers.Validator(payload); err != nil {
		log.Println("Error validation ", err)
		res.Code = http.StatusBadRequest
		res.Message = "Validation Error"

		ctx.JSON(res.Code, res)
		return
	}
	res = c.Usecase.CreateCoupon(ctx, payload)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetListCoupon(ctx *gin.Context) {
	log.Println("<<<<GetListCouponControllers>>>")

	var res models.Response
	var params models.GetListCouponParams

	if err := ctx.BindQuery(&params); err != nil {
		log.Println("Error bad request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.GetListCoupon(ctx, params)

	ctx.JSON(res.Code, res)
}
