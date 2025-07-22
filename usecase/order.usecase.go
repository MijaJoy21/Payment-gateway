package usecase

import (
	"log"
	"math"
	"net/http"
	"os"
	"payment-gateway/helpers"
	"payment-gateway/models"
	"payment-gateway/repository/entity"
	"strings"

	"github.com/gin-gonic/gin"
)

func (u *usecase) CreateOrder(ctx *gin.Context, payload models.ReqCreateOrder) models.Response {
	res := models.Response{}

	userData := &models.ClaimsJwt{}

	if ctx.Value("user") != nil {
		userData = ctx.Value("user").(*models.ClaimsJwt)
	}
	statusConfirmed := 0
	order := &entity.Order{
		UserId:       userData.Id,
		ExpeditionId: 1,
		InvoiceId:    helpers.GenerateInvoiceId(),
		StatusOrder:  &statusConfirmed,
		CouponId:     payload.CouponId,
	}

	if err := u.Repository.CreateOrder(ctx, order); err != nil {
		log.Println("Error create order", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"
		return res
	}

	orderDetail := []entity.OrderDetail{}
	for _, val := range payload.OrderDetail {
		orderDetail = append(orderDetail, entity.OrderDetail{
			OrderId:   order.Id,
			ProductId: val.ProductId,
			Quantity:  val.Quantity,
		})
	}

	if err := u.Repository.CreateOrderDetail(ctx, orderDetail); err != nil {
		log.Println("Error create order detail", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"
		return res
	}

	productIds := []int{}
	for _, val := range payload.OrderDetail {
		productIds = append(productIds, val.ProductId)
	}
	if err := u.Repository.DeleteCartByUserIdAndProductId(ctx, userData.Id, productIds); err != nil {
		log.Println("error delete")
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}

func (u *usecase) GetAllOrder(ctx *gin.Context, params models.GetAllOrderParams) models.Response {
	res := models.Response{}

	data, total, err := u.Repository.GetAllOrder(ctx, params)

	if err != nil {
		log.Println("Error Get All Order", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Error Get All Order"

		return res
	}
	response := []models.GetAllOrderRes{}

	for _, val := range data {
		response = append(response, models.GetAllOrderRes{
			Id:        val.Id,
			InvoiceID: val.InvoiceId,
			Status:    val.StatusOrderName(),
			UserName:  val.User.Name,
			UserEmail: val.User.Email,
		})
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = response
	res.Pagination = &models.Pagination{
		Page:      params.Page,
		Limit:     params.Limit,
		TotalData: total,
		LastPage:  int(math.Ceil(float64(total) / float64(params.Limit))),
	}
	return res
}

func (u *usecase) UpdateOrderStatusById(ctx *gin.Context, id int, payload models.UpdateOrderStatusByIdReq) models.Response {
	res := models.Response{}

	if err := u.Repository.UpdateOrderStatusById(ctx, payload.Status, id); err != nil {
		log.Println("Error update order status by id", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Unprocessable Entity"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	return res
}

func (u *usecase) GetOrderById(ctx *gin.Context, id int) models.Response {
	res := models.Response{}

	data, err := u.Repository.GetOrderById(ctx, id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data not found"

		return res
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	return res
}

func (u *usecase) GetHistoryOrderByUserId(ctx *gin.Context, userId int, params models.GetAllHistoryOrderParams) models.Response {
	res := models.Response{}

	data, total, err := u.Repository.GetHistoryOrderByUserId(ctx, userId, params)

	if err != nil {
		log.Println("Error get history order by user id", err)
		res.Code = http.StatusOK
		res.Message = "Internal Server Error"

		return res
	}

	response := []models.GetAllHistoryOrderRes{}
	for _, val := range data {
		for _, orderDetailVal := range val.OrderDetail {
			response = append(response, models.GetAllHistoryOrderRes{
				Id:           orderDetailVal.Id,
				InvoiceId:    val.InvoiceId,
				ProductName:  orderDetailVal.Product.Name,
				ProductImage: os.Getenv("ADDRESS_SERVICE") + os.Getenv("PORT") + strings.Split(orderDetailVal.Product.Image, ",")[0],
				Status:       val.StatusOrderName(),
				Quantity:     orderDetailVal.Quantity,
				Price:        orderDetailVal.Product.Price,
				CreatedAt:    val.CreatedAt,
			})
		}
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = response
	res.Pagination = &models.Pagination{
		Page:      params.Page,
		Limit:     params.Limit,
		TotalData: total,
		LastPage:  int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return res
}
