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
	}
)

func InitControllers(usecase usecase.Usecase) Controllers {
	return &controllers{
		Usecase: usecase,
	}
}
