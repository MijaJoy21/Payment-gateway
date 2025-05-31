package usecase

import (
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"payment-gateway/models"
	"payment-gateway/repository/entity"
	"strings"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateReview(ctx *gin.Context, file []*multipart.FileHeader, payload models.CreateReview) models.Response {
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

	data := entity.Review{
		Description: payload.Description,
		IsAnon:      payload.IsAnon,
		Image:       strings.Join(fileNames, ","),
	}

	if err := uc.Repository.CreateReview(ctx, &data); err != nil {
		log.Println("Error Create Review", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetReviewById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetReviewById(ctx, id)

	if err != nil {
		log.Println("Eror Data Not Found")
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (uc *usecase) PutReview(ctx *gin.Context, id int, payload models.RequestReview) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetReviewById(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Review{
		Description: payload.Description,
		IsAnon:      payload.IsAnon,
	}

	if err := uc.Repository.PutReview(ctx, id, updatedData); err != nil {
		log.Println("Error updating Review :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed to update review"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}

func (uc *usecase) DeleteReview(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	if err := uc.Repository.DeleteReview(ctx, id); err != nil {
		log.Println("Error Data not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Review Deleted"
	return res
}
