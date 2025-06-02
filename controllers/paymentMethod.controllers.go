package controllers

import (
	"log"
	"net/http"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreatePaymentMethod(ctx *gin.Context) {
	log.Println("<<Controllers Create Payment Method>>")
	var res models.Response
	payload := models.ReqPaymentMethod{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind Json", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreatePaymentMethod(ctx, payload)
	log.Println("Response Create Payment Method", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetAllPaymentMethods(ctx *gin.Context) {
	res := c.Usecase.GetAllPaymentMethods(ctx)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetPaymentMethodById(ctx *gin.Context) {
	log.Println("<<<Controllers Get user by ID>>>")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetPaymentMethodById(ctx, id)
	log.Println("Response Get Detail Paymeent Method By id", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutPaymentMethod(ctx *gin.Context) {
	log.Println("<<Controllers Update Payment Method>>")
	var res models.Response

	// Ambil ID dari parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID parameter"
		ctx.JSON(res.Code, res)
		return
	}

	// Bind JSON ke payload
	payload := models.ReqPaymentMethod{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request body"
		ctx.JSON(res.Code, res)
		return
	}

	// Panggil usecase untuk memperbarui data
	res = c.Usecase.PutPaymentMethod(ctx, id, payload)
	log.Println("Response Update Payment Method", res)

	ctx.JSON(res.Code, res)
}
