package routes

import (
	"fmt"
	"payment-gateway/controllers"
	"payment-gateway/helpers"
	"payment-gateway/middleware"

	"github.com/gin-gonic/gin"
)

type (
	Router struct {
		controllers controllers.Controllers
		gin         *gin.Engine
	}

	RouterInterface interface {
		StartGinServer() error
	}
)

func InitRoutes(ctrl controllers.Controllers) RouterInterface {
	return &Router{
		controllers: ctrl,
		gin:         gin.New(),
	}
}

func (r *Router) StartGinServer() error {
	fmt.Println("Start Server")
	//prefix api
	api := r.gin.Group("/api")
	api.GET("/hc", r.controllers.GetHealthCheck)
	api.GET("/user", r.controllers.GetAllUsers)
	auth := api.Group("/authorization")
	{
		auth.POST("/registration", r.controllers.RegistrationUser)
		auth.POST("/login", r.controllers.LoginUser)
	}

	user := api.Group("/user")
	{
		user.GET("/", middleware.Authorization(""), r.controllers.GetUserById)
		user.POST("/register/admin", r.controllers.RegistrationAdmin)

	}

	payment := api.Group("/payment")
	{
		payment.POST("/method", middleware.Authorization("Admin"), r.controllers.CreatePaymentMethod)
		payment.GET("/payment-method", middleware.Authorization("Admin"), r.controllers.GetAllPaymentMethods)
		payment.POST("/pay", r.controllers.CreatePayment)
		payment.GET("/:id", r.controllers.GetPaymentMethodById)
		payment.PUT("/method/:id", middleware.Authorization("Admin"), r.controllers.PutPaymentMethod)
	}

	category := api.Group("/category")
	{
		category.POST("/create", middleware.Authorization("Admin"), r.controllers.CreateCategory)
		category.GET("/", r.controllers.GetAllCategory)
		category.GET("/:id", r.controllers.GetCategoryById)
		category.PUT("update/:id", middleware.Authorization("Admin"), r.controllers.PutCategory)
	}

	etalase := api.Group("/etalase")
	{
		etalase.POST("/create", middleware.Authorization("Admin"), r.controllers.CreateEtalase)
		etalase.GET("/", r.controllers.GetAllEtalase)
		etalase.GET("/:id", r.controllers.GetEtalaseById)
		etalase.PUT("/update/:id", middleware.Authorization("Admin"), r.controllers.PutEtalase)
	}

	expedition := api.Group("/expedition")
	{
		expedition.POST("/create", middleware.Authorization("Admin"), r.controllers.CreateExpedition)
		expedition.GET("/", r.controllers.GetAllExpedition)
		expedition.GET("/:id", r.controllers.GetExpeditionById)
		expedition.PUT("/update/:id", middleware.Authorization("Admin"), r.controllers.PutExpediton)
	}

	product := api.Group("/product")
	{
		product.POST("/create", middleware.Authorization("Admin"), r.controllers.CreateProduct)
	}

	if err := helpers.StartGinServer(r.gin); err != nil {
		fmt.Println("Error Start Server", err)
	}

	return nil
}
