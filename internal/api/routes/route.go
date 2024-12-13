package routes

import (
	"Marketplace/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(
	router *gin.RouterGroup,
	companyHandler *handler.CompanyHandler,
	influencerHandler *handler.InfluencerHandler,
	userHandler *handler.UserHandler,
	adHandler *handler.AdHandler,
	orderHandler *handler.OrderHandler,
	applicationHandler *handler.ApplicationHandler,
) {

	user := router.Group("/user")
	{
		user.GET("/account_type/:id", userHandler.GetAccountType)
	}

	company := router.Group("/company")
	{
		company.POST("/login", companyHandler.Login)
		company.POST("/register", companyHandler.Register)
		company.POST("/verify", companyHandler.VerifyEmail)
		company.POST("/verify/send_code", companyHandler.SendCode)
		company.GET("/:id", companyHandler.GetByID)
		company.GET("/email/:email", companyHandler.GetByEmail)
	}

	influencer := router.Group("/influencer")
	{
		influencer.POST("/login", influencerHandler.Login)
		influencer.POST("/register", influencerHandler.Register)
		influencer.POST("/verify", influencerHandler.VerifyEmail)
		influencer.POST("/verify/send_code", influencerHandler.SendCode)
		influencer.GET("/:id", companyHandler.GetByID)
		influencer.GET("email/:email", companyHandler.GetByEmail)
	}

	ad := router.Group("/ad").Use(JWTMiddleware())
	{
		ad.POST("/create", adHandler.Create)
		ad.GET("/:id", adHandler.GetByID)
		ad.GET("/all", adHandler.GetAll)
		ad.PUT("/:id/status", adHandler.UpdateStatus)
		ad.GET("/company/:id", adHandler.GetByCompanyID)
		ad.DELETE("/delete/:id", adHandler.Delete)
	}

	order := router.Group("/order").Use(JWTMiddleware())
	{
		order.POST("/create", orderHandler.Create)
		order.GET("/:id", orderHandler.GetByID)
		order.PUT("/:id/status", orderHandler.UpdateStatus)
		order.GET("/company/:id", orderHandler.GetByCompanyID)
		order.GET("/influencer/:id", orderHandler.GetByInfluencerID)
		order.DELETE("/:id/delete", orderHandler.Delete)
	}

	application := router.Group("/application").Use(JWTMiddleware())
	{
		application.POST("/create", applicationHandler.Create)
		application.GET("/:id", applicationHandler.GetByID)
		application.GET("/ad/:id", applicationHandler.GetByAdID)
		application.GET("/influencer/:id", applicationHandler.GetByInfluencer)
		application.PUT("/:id/status", applicationHandler.UpdateStatus)
		application.DELETE("/:id/delete", applicationHandler.Delete)
	}

	//notification := router.Group("/notification")
	//{
	//	notification.POST("/send", notificationHandler.Send)
	//	notification.GET("/user/:id", notificationHandler.GetByUserID)
	//	notification.PUT("/:id/mark_as_read", NotificationHandler.MarkAsRead)
	//}

	//moderator := router.Group("/moderator)
	//{
	//	moderator.GET("/moderator/:list
}
