package routes

import (
	"fmt"
	"os"
	"payment-gateway/controllers"
	"payment-gateway/helpers"
	"payment-gateway/middleware"
	"time"

	"github.com/gin-contrib/cors"
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
	//cors
	r.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"}, // frontend kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Prefix Api
	r.gin.Static("/uploads", os.Getenv("IMAGE_UPLOAD"))
	//prefix api
	r.gin.Static("/image", os.Getenv("IMAGE_UPLOAD"))
	api := r.gin.Group("/api")
	api.GET("/hc", r.controllers.GetHealthCheck)
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
		product.GET("/", r.controllers.GetAllProduct)
		product.GET("/:id", r.controllers.GetProductById)
		product.PUT("/update/:id", middleware.Authorization("Admin"), r.controllers.PutProduct)
		product.DELETE("delete/:id", middleware.Authorization("Admin"), r.controllers.DeleteProduct)
	}

	cart := api.Group("/cart")
	{
		cart.POST("/create", middleware.Authorization(""), r.controllers.CreateCart)
		cart.GET("/:id", middleware.Authorization(""), r.controllers.GetCartById)
		cart.PUT("/update/:id", r.controllers.PutCart)
		cart.DELETE("/delete/:id", r.controllers.DeleteCart)
	}
	review := api.Group("/review")
	{
		review.POST("/create", r.controllers.CreateReview)
		review.GET("/:id", r.controllers.GetReviewById)
		review.PUT("update/:id", r.controllers.PutReview)
		review.DELETE("delete/:id", r.controllers.DeleteReview)
	}

	order := api.Group("/order")
	{
		order.POST("/create", middleware.Authorization(""), r.controllers.CreateOrder)
		order.GET("/all/admin", middleware.Authorization("Admin"), r.controllers.GetAllOrder)
		order.PUT("/update-status/:id", middleware.Authorization("Admin"), r.controllers.UpdateOrderStatusById)
		order.GET("/detail/admin/:id", middleware.Authorization("Admin"), r.controllers.GetOrderById)
	}

	if err := helpers.StartGinServer(r.gin); err != nil {
		fmt.Println("Error Start Server", err)
	}

	return nil
}
