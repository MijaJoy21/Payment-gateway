package usecase

import (
	"log"
	"net/http"
	"os"
	"payment-gateway/models"
	"payment-gateway/repository/entity"
	"strings"

	"github.com/gin-gonic/gin"
)

func (uc *usecase) CreateCart(ctx *gin.Context, payload models.CreateCart) models.Response {
	res := models.Response{}
	data := entity.Cart{
		UserId:    payload.UserId,
		ProductId: payload.ProductId,
		Quantity:  payload.Quantity,
	}

	if err := uc.Repository.CreateCart(ctx, data); err != nil {
		log.Println("Error Create Cart", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"

	return res
}

func (uc *usecase) GetCartById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := uc.Repository.GetCartByid(ctx, id)

	if err != nil {
		log.Println("Error Data Not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"

		return res
	}

	cart := []models.ResponseListCart{}

	for _, val := range data {
		cart = append(cart, models.ResponseListCart{
			Id:              val.Id,
			ProductId:       val.ProductId,
			ProductName:     val.Product.Name,
			ProductQuantity: val.Quantity,
			ProductImage:    os.Getenv("ADDRESS_SERVICE") + os.Getenv("PORT") + strings.Split(val.Product.Image, ",")[0],
			ProductPrice:    val.Product.Price,
		})
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = cart

	return res
}

func (uc *usecase) PutCart(ctx *gin.Context, id int, payload models.RequestCart) models.Response {
	res := models.Response{}

	_, err := uc.Repository.GetCartByid(ctx, id)
	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		return res
	}

	updatedData := entity.Cart{
		Quantity: payload.Quantity,
	}

	if payload.Quantity < 1 {
		res.Code = http.StatusBadRequest
		res.Message = "Quanttiy must be at least 1"
	}

	if err := uc.Repository.PutCart(ctx, id, updatedData); err != nil {
		log.Println("Error Updating Category :", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Failed To update Cart"
		return res
	}

	res.Code = http.StatusOK
	res.Message = "Cart updated successfully"
	return res
}

func (uc *usecase) DeleteCart(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	if err := uc.Repository.DeleteCart(ctx, id); err != nil {
		log.Println("Error Data Not Found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data not Found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Cart Deleted"
	return res
}
