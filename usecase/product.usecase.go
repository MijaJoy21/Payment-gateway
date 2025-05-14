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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateProduct(ctx *gin.Context, file []*multipart.FileHeader, payload models.CreateProduct) models.Response {
	res := models.Response{}
	fileNames := []string{}
	for _, val := range file {
		filePath := filepath.Join(os.Getenv("IMAGE_UPLOAD"), val.Filename)
		filePath = filepath.ToSlash(filePath)

		if _, err := os.Stat(os.Getenv("IMAGE_UPLOAD")); os.IsNotExist(err) {
			os.Mkdir(os.Getenv("IMAGE_UPLOAD"), os.ModePerm)
		}

		if err := ctx.SaveUploadedFile(val, filePath); err != nil {
			log.Println("Error upload image ", err)
			res.Code = http.StatusUnprocessableEntity
			res.Message = "Failed Upload Image"

			return res
		}

		fileNames = append(fileNames, filePath)
	}
	tmpTime := 0
	if payload.PreOrderDate != "" {
		var err error
		tmpTime, err = strconv.Atoi(payload.PreOrderDate)

		if err != nil {
			log.Println("Error Convert Time ", err)
			res.Code = http.StatusBadRequest
			res.Message = "Error time date format"

			return res
		}
	}

	preOrderTime := time.Unix(int64(tmpTime), 0)

	data := entity.Product{
		Name:        payload.Name,
		Categoryid:  payload.CategoryId,
		Description: payload.Description,
		Price:       payload.Price,
		Image:       strings.Join(fileNames, ","),
		Status:      payload.Status,
		Weight:      payload.Weight,
	}

	if payload.IsPreOrder == 1 {
		data.IsPreorder = payload.IsPreOrder
		data.PreOrderDate = &preOrderTime
	}

	if err := uc.Repository.CreateProduct(ctx, &data); err != nil {
		log.Println("Error Create Product", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	if len(payload.Variant) != 0 {
		variantProduct := []entity.Variant{}
		for _, val := range payload.Variant {
			variantProduct = append(variantProduct, entity.Variant{
				ProductId: data.Id,
				Name:      val.Name,
				Price:     val.Price,
				Weight:    val.Weight,
				Quantity:  val.Quantity,
				Status:    val.Status,
			})
		}

		if err := uc.Repository.CreateVariant(ctx, variantProduct); err != nil {
			log.Println("Error create product", err)
			res.Code = http.StatusUnprocessableEntity
			res.Message = "Unprocessable entity"

			return res
		}
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

	images := []string{}
	for _, val := range strings.Split(data.Image, ",") {
		tmp := os.Getenv("ADDRESS_SERVICE") + os.Getenv("PORT") + val
		images = append(images, tmp)
	}

	response := models.GetDetailProductResponse{
		Id:       data.Id,
		Name:     data.Name,
		Price:    data.Price,
		Weight:   data.Weight,
		Quantity: data.Quantity,
		Image:    images,
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = response

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

func (uc *usecase) DeleteProduct(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	if err := uc.Repository.DeleteProduct(ctx, id); err != nil {
		log.Println("Error Data not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Product Deleted"
	return res
}
