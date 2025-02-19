package controllers

import (
	"log"
	"net/http"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateCategory(ctx *gin.Context) {
	log.Println("Controllers Registration Category")
	var res models.Response
	payload := models.RequestCategory{}

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

	res = c.Usecase.CreateCategory(ctx, payload)
	log.Println("Response Registration User", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetAllCategory(ctx *gin.Context) {
	var res models.Response

	res = c.Usecase.GetAllCategory(ctx)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetCategoryById(ctx *gin.Context) {
	log.Println("Controllers Get Category By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetCategoryById(ctx, id)
	log.Println("Response Get Detail Category By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutCategory(ctx *gin.Context) {
	log.Println("Controllers Update Category")
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID Parameter"
		ctx.JSON(res.Code, res)
		return
	}

	payload := models.RequestCategory{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request body"
		ctx.JSON(res.Code, res)
		return
	}

	err = helpers.Validator(payload)
	if err != nil {
		log.Println("Error Update Field Empty", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please Filled The required field"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutCategory(ctx, id, payload)
	log.Println("Response Update Category", res)

	ctx.JSON(res.Code, res)
}
