package controllers

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateCart(ctx *gin.Context) {
	log.Println("Controllers Create Cart")
	var res models.Response
	payload := models.CreateCart{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateCart(ctx, payload)
	log.Println("Response Create Cart", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetCartById(ctx *gin.Context) {
	log.Println("Controllers Get Cart By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetCartById(ctx, id)
	log.Println("Response Get Detail Cart By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutCart(ctx *gin.Context) {
	log.Println("Controllers Update Product")
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID Parameter"
		ctx.JSON(res.Code, res)
		return
	}

	payload := models.RequestCart{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request Body"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutCart(ctx, id, payload)
	log.Println("Response Update Cart", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) DeleteCart(ctx *gin.Context) {
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID"
		return
	}

	res = c.Usecase.DeleteCart(ctx, id)

	ctx.JSON(res.Code, res)
}
