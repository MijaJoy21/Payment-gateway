package controllers

import (
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateEtalase(ctx *gin.Context) {
	log.Println("Controllers Registration Etalase")
	var res models.Response
	payload := models.RequestEtalase{}

	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	err := helpers.Validator(payload)
	if err != nil {
		log.Println("Error", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please Filled The required field"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateEtalase(ctx, payload)
	log.Println("Response Registration Etalase")

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetAllEtalase(ctx *gin.Context) {
	var res models.Response

	res = c.Usecase.GetAllEtalase(ctx)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetEtalaseById(ctx *gin.Context) {
	log.Println("Controllers Get Etalase By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetEtalaseById(ctx, id)
	log.Println("Response Get Detail Etalse By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutEtalase(ctx *gin.Context) {
	log.Println("Controllers Update Etalase")
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID Parameter"
		ctx.JSON(res.Code, res)
		return
	}

	payload := models.RequestEtalase{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error Binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request body"
		ctx.JSON(res.Code, res)
		return
	}

	if err != nil {
		log.Println("Error", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please Filled The required field"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutEtalase(ctx, id, payload)
	log.Println("Response Update Etalase", res)

	ctx.JSON(res.Code, res)

}
