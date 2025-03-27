package usecase

import (
	"log"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"payment-gateway/models"
	"payment-gateway/repository/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateProduct(ctx *gin.Context, file *multipart.FileHeader, payload models.CreateProduct) models.Response {
	res := models.Response{}
	filePath := filepath.Join(os.Getenv("IMAGE_UPLOAD"), file.Filename)
	filePath = filepath.ToSlash(filePath)

	if _, err := os.Stat(os.Getenv("IMAGE_UPLOAD")); os.IsNotExist(err) {
		os.Mkdir(os.Getenv("IMAGE_UPLOAD"), os.ModePerm)
	}

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error upload image ", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Upload Image"

		return res
	}

	tmpTime, err := strconv.Atoi(payload.PreOrderDate)

	if err != nil {
		log.Println("Error Convert Time ", err)
		res.Code = http.StatusBadRequest
		res.Message = "Error time date format"

		return res
	}

	preOrderTime := time.Unix(int64(tmpTime), 0)

	data := entity.Product{
		Name:        payload.Name,
		Categoryid:  payload.CategoryId,
		Description: payload.Description,
		Price:       payload.Price,
		Image:       file.Filename,
		Status:      payload.Status,
		Weight:      payload.Weight,
	}

	if payload.IsPreOrder == 1 {
		data.IsPreorder = payload.IsPreOrder
		data.PreOrderDate = &preOrderTime
	}

	if err := uc.Repository.CreateProduct(ctx, data); err != nil {
		log.Println("Error Create Product", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetAllProduct(ctx *gin.Context, params models.ParamsGetProduct) models.Response {
	res := models.Response{}

	data, total, err := uc.Repository.GetProduct(ctx, params)

	if err != nil {
		log.Println("Error Get All Product", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

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

func (uc *usecase) GetProductById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetProductById(ctx, id)

	if err != nil {
		log.Println("Error Data Not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) PutProduct(ctx *gin.Context, id int, payload models.RequestProduct) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetProductById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Product{
		Name:        payload.Name,
		Categoryid:  payload.CategoryId,
		Description: payload.Description,
		Price:       payload.Price,
		Status:      payload.Status,
		IsPreorder:  payload.IsPreOrder,
		Weight:      payload.Weight,
	}

	if err := uc.Repository.PutProduct(ctx, id, updatedData); err != nil {
		log.Println("Error updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update Category"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}
