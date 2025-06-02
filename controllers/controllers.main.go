package controllers

import (
	"payment-gateway/usecase"

	"github.com/gin-gonic/gin"
)

type (
	controllers struct {
		Usecase usecase.Usecase
	}

	Controllers interface {
		GetHealthCheck(ctx *gin.Context)
		GetAllUsers(ctx *gin.Context)
		RegistrationUser(ctx *gin.Context)
		GetUserById(ctx *gin.Context)
		LoginUser(ctx *gin.Context)
		CreatePaymentMethod(ctx *gin.Context)
		GetAllPaymentMethods(ctx *gin.Context)
		CreatePayment(ctx *gin.Context)
		GetPaymentMethodById(ctx *gin.Context)
		PutPaymentMethod(ctx *gin.Context)
		RegistrationAdmin(ctx *gin.Context)
		CreateCategory(ctx *gin.Context)
		GetAllCategory(ctx *gin.Context)
		GetCategoryById(ctx *gin.Context)
		PutCategory(ctx *gin.Context)
		CreateEtalase(ctx *gin.Context)
		GetAllEtalase(ctx *gin.Context)
		GetEtalaseById(ctx *gin.Context)
		PutEtalase(ctx *gin.Context)
		CreateExpedition(ctx *gin.Context)
		GetAllExpedition(ctx *gin.Context)
		GetExpeditionById(ctx *gin.Context)
		PutExpediton(ctx *gin.Context)
		CreateProduct(ctx *gin.Context)
		GetAllProduct(ctx *gin.Context)
		GetProductById(ctx *gin.Context)
		PutProduct(ctx *gin.Context)
		CreateCart(ctx *gin.Context)
		GetCartById(ctx *gin.Context)
		PutCart(ctx *gin.Context)
		DeleteCart(ctx *gin.Context)
		DeleteProduct(ctx *gin.Context)
		CreateOrder(ctx *gin.Context)
		GetAllOrder(ctx *gin.Context)
		UpdateOrderStatusById(ctx *gin.Context)
		GetOrderById(ctx *gin.Context)
		CreateReview(ctx *gin.Context)
		GetReviewById(ctx *gin.Context)
		PutReview(ctx *gin.Context)
		DeleteReview(ctx *gin.Context)
	}
)

func InitControllers(usecase usecase.Usecase) Controllers {
	return &controllers{
		Usecase: usecase,
	}
}
