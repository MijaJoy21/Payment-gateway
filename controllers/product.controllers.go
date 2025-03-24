package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"payment-gateway/constant"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateProduct(ctx *gin.Context) {
	log.Println("<<Controllers Create Product>>")
	var res models.Response
	payload := models.CreateProduct{}

	file, err := ctx.FormFile("image")

	if err != nil {
		log.Println("Error get image", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	ext := filepath.Ext(file.Filename)

	if !constant.AllowedExtensions[ext] {
		log.Println("Image type not supported", err)
		res.Code = http.StatusBadRequest
		res.Message = "Image file not supported"

		ctx.JSON(res.Code, res)
		return
	}

	if err := ctx.ShouldBind(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateProduct(ctx, file, payload)
	log.Println("Response Create Product", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetAllProduct(ctx *gin.Context) {
	var res models.Response
	params := models.ParamsGetProduct{}

	if err := ctx.BindQuery(&params); err != nil {
		log.Println("Error Bind Params", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.GetAllProduct(ctx, params)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetProductById(ctx *gin.Context) {
	log.Println("Controllers Get Product By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetProductById(ctx, id)
	log.Println("Response Get Detail Category By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutProduct(ctx *gin.Context) {
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

	payload := models.RequestProduct{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request body"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutProduct(ctx, id, payload)
	log.Println("Response Update Product", res)

	ctx.JSON(res.Code, res)
}
