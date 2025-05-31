package controllers

import (
	"log"
	"net/http"
	"path/filepath"
	"payment-gateway/constant"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (c *controllers) CreateReview(ctx *gin.Context) {
	log.Println("<<Controllers Create Review>>")
	var res models.Response
	payload := models.CreateReview{}

	form, err := ctx.MultipartForm()

	if err != nil {
		log.Println("Error get image", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	files := form.File["image"]
	if len(files) == 0 {
		log.Println("Image required")
		res.Code = http.StatusBadRequest
		res.Message = "Image required"

		ctx.JSON(res.Code, res)
		return
	}

	for _, val := range files {
		ext := filepath.Ext(val.Filename)
		if !constant.AllowedExtensions[ext] {
			log.Println("Image type not supported", err)
			res.Code = http.StatusBadRequest
			res.Message = "Image file not supported"

			ctx.JSON(res.Code, res)
			return
		}
	}

	if err := ctx.ShouldBind(&payload); err != nil {
		log.Println("Error Bind JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helpers.Validator(payload); err != nil {
		log.Println("Error Validation", err)
		res.Code = http.StatusBadRequest
		res.Message = "Please Filled required filled"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.CreateReview(ctx, files, payload)
	log.Println("Response Create Review", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) GetReviewById(ctx *gin.Context) {
	log.Println("Controllers Get Review By ID")
	var res models.Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	res = c.Usecase.GetReviewById(ctx, id)
	log.Println("Response Get Detail Review By ID", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) PutReview(ctx *gin.Context) {
	log.Println("Controllers Update Review")
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println("Invalid ID parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID Parameter"
		ctx.JSON(res.Code, res)
		return
	}

	payload := models.RequestReview{}
	if err := ctx.BindJSON(&payload); err != nil {
		log.Println("Error binding JSON", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid request body"
		ctx.JSON(res.Code, res)
		return
	}

	res = c.Usecase.PutReview(ctx, id, payload)
	log.Println("Response Update Review", res)

	ctx.JSON(res.Code, res)
}

func (c *controllers) DeleteReview(ctx *gin.Context) {
	var res models.Response

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id <= 0 {
		log.Println("Invalid ID Parameter", err)
		res.Code = http.StatusBadRequest
		res.Message = "Invalid ID"
		return
	}

	res = c.Usecase.DeleteReview(ctx, id)
	ctx.JSON(res.Code, res)
}
