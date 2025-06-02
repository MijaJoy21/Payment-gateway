package controllers

import (
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateOrder(ctx *gin.Context) {
	log.Println("<<Controllers Create Order>>")
	var res models.Response
	payload := models.ReqCreateOrder{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind Json", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helpers.Validator(payload); err != nil {
		log.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateOrder(ctx, payload)
	log.Println("Response Create Order", res)

	ctx.JSON(res.Code, res)

}

func (c *controllers) GetAllOrder(ctx *gin.Context) {
	log.Println("<<Controllers Create Order>>")
	var res models.Response
	var params models.GetAllOrderParams

	if err := ctx.BindQuery(&params); err != nil {
		log.Println("Error bad request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helpers.Validator(params); err != nil {
		log.Println("Error validate", err)
		res.Code = http.StatusBadRequest
		res.Message = "Validation Error"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.GetAllOrder(ctx, params)

	ctx.JSON(res.Code, res)
}

func (c *controllers) UpdateOrderStatusById(ctx *gin.Context) {
	log.Println("<<Controllers Create Order>>")
	var res models.Response
	payload := models.UpdateOrderStatusByIdReq{}

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.UpdateOrderStatusById(ctx, id, payload)

	ctx.JSON(res.Code, res)
}
