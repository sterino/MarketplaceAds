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
	adHandler *handler.AdHandler, // Добавлен хендлер для объявлений
) {

	// Routes for user
	user := router.Group("/user")
	{
		user.GET("/account_type/:id", userHandler.GetAccountType)
		// Add other user routes as needed
	}

	// Routes for company
	company := router.Group("/company")
	{
		company.POST("/login", companyHandler.Login)
		company.POST("/register", companyHandler.Register)
		company.POST("/verify", companyHandler.VerifyEmail)
		company.POST("/send_code", companyHandler.SendCode)
		// Add other company routes as needed
	}

	// Routes for influencer
	influencer := router.Group("/influencer")
	{
		influencer.POST("/login", influencerHandler.Login)
		influencer.POST("/register", influencerHandler.Register)
		// Add other influencer routes as needed
	}

	// Routes for ads
	ad := router.Group("/ad")
	{
		ad.POST("/create", adHandler.Create)          // Создание объявления
		ad.GET("/:id", adHandler.GetByID)             // Получение объявления по ID
		ad.GET("/all", adHandler.GetAll)              // Получение всех объявлений
		ad.PUT("/:id/status", adHandler.UpdateStatus) // Обновление статуса объявления
		// Add other ad-related routes as needed
	}

	// Add other routes as needed
}
