package usecase

import (
	"mime/multipart"
	"payment-gateway/models"
	"payment-gateway/repository"

	"github.com/gin-gonic/gin"
)

type (
	usecase struct {
		Repository repository.Repository
	}

	Usecase interface {
		GetHealthCheck(ctx *gin.Context) models.Response
		GetAllUsers(ctx *gin.Context) models.Response
		RegistrationUser(ctx *gin.Context, payload models.ReqRegistrationUser) models.Response
		GetUserById(ctx *gin.Context) models.Response
		LoginUser(ctx *gin.Context, payload models.RegLogin) models.Response
		CreatePaymentMethod(ctx *gin.Context, payload models.ReqPaymentMethod) models.Response
		GetAllPaymentMethods(ctx *gin.Context) models.Response
		CreatePayment(ctx *gin.Context, payload models.ReqPayment) models.Response
		GetPaymentMethodById(ctx *gin.Context, id int) models.Response
		PutPaymentMethod(ctx *gin.Context, id int, payload models.ReqPaymentMethod) models.Response
		RegistrationAdmin(ctx *gin.Context, payload models.ReqRegistrationUser) models.Response
		CreateCategory(ctx *gin.Context, payload models.RequestCategory) models.Response
		GetAllCategory(ctx *gin.Context) models.Response
		GetCategoryById(ctx *gin.Context, id int) models.Response
		PutCategory(ctx *gin.Context, id int, payload models.RequestCategory) models.Response
		CreateEtalase(ctx *gin.Context, payload models.RequestEtalase) models.Response
		GetAllEtalase(ctx *gin.Context) models.Response
		GetEtalaseById(ctx *gin.Context, id int) models.Response
		PutEtalase(ctx *gin.Context, id int, payload models.RequestEtalase) models.Response
		CreateExpedition(ctx *gin.Context, payload models.RequestExpedition) models.Response
		GetAllExpedition(ctx *gin.Context, params models.ParamsGetExpeditions) models.Response
		GetExpeditionById(ctx *gin.Context, id int) models.Response
		PutExpedition(ctx *gin.Context, id int, payload models.RequestExpedition) models.Response
		CreateProduct(ctx *gin.Context, file []*multipart.FileHeader, payload models.CreateProduct) models.Response
		GetAllProduct(ctx *gin.Context, params models.ParamsGetProduct) models.Response
		GetProductById(ctx *gin.Context, id int) models.Response
		PutProduct(ctx *gin.Context, id int, payload models.RequestProduct) models.Response
		CreateCart(ctx *gin.Context, payload models.CreateCart) models.Response
		GetCartById(ctx *gin.Context, id int) models.Response
		PutCart(ctx *gin.Context, id int, payload models.RequestCart) models.Response
		DeleteCart(ctx *gin.Context, id int) models.Response
		DeleteProduct(ctx *gin.Context, id int) models.Response
		CreateOrder(ctx *gin.Context, payload models.ReqCreateOrder) models.Response
		GetAllOrder(ctx *gin.Context, params models.GetAllOrderParams) models.Response
		UpdateOrderStatusById(ctx *gin.Context, id int, payload models.UpdateOrderStatusByIdReq) models.Response
		GetOrderById(ctx *gin.Context, id int) models.Response
		CreateReview(ctx *gin.Context, file []*multipart.FileHeader, payload models.CreateReview) models.Response
		GetReviewById(ctx *gin.Context, id int) models.Response
		PutReview(ctx *gin.Context, id int, payload models.RequestReview) models.Response
		DeleteReview(ctx *gin.Context, id int) models.Response
	}
)

func InitUsecase(repository repository.Repository) Usecase {
	return &usecase{
		Repository: repository,
	}

}
