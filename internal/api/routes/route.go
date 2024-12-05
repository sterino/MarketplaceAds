package routes

import (
	"Marketplace/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, companyHandler *handler.CompanyHandler, influencerHandler *handler.InfluencerHandler) {
	company := router.Group("/company")
	{
		//company.GET("/")
		company.POST("/login", companyHandler.Login)
		company.POST("/register", companyHandler.Register)
		company.POST("/verify", companyHandler.VerifyEmail)
		//company.PUT("/")
		//company.GET("/")
	}

	influencer := router.Group("/")
	{
		influencer.POST("/login", influencerHandler.Login)
		influencer.POST("/register", influencerHandler.Register)
		//influencer.GET("/")
		//influencer.POST("/")
		//influencer.PUT("/")
	}

}
