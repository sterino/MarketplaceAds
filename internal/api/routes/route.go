package routes

import (
	"Marketplace/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.RouterGroup, companyHandler *handler.CompanyHandler) {
	company := router.Group("/company")
	{
		//company.GET("/")
		company.POST("/login", companyHandler.Login)
		company.POST("/register", companyHandler.Register)
		//company.PUT("/")
		//company.GET("/")
	}

	//influencer := router.Group("/")
	//{
	//	influencer.GET("/")
	//	influencer.POST("/")
	//	influencer.PUT("/")
	//}

}
