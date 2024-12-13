package api

import (
	"Marketplace/internal/api/handler"
	"Marketplace/internal/api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"time"
)

type Server struct {
	engine *gin.Engine
}

func NewServer(
	companyHandler *handler.CompanyHandler,
	influenceHandler *handler.InfluencerHandler,
	userHandler *handler.UserHandler,
	adHandler *handler.AdHandler,
	orderHandler *handler.OrderHandler,
	applicationHandler *handler.ApplicationHandler,
) *Server {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(MethodNotAllowedMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.InitRoutes(router.Group("/api"), companyHandler, influenceHandler, userHandler, adHandler, orderHandler, applicationHandler)

	return &Server{router}
}

func (s *Server) Run(infoLog *log.Logger, errorLog *log.Logger) {
	infoLog.Printf("starting server on: 8000")
	err := s.engine.Run(":8000")
	errorLog.Fatal(err)
}

func MethodNotAllowedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		allowedMethods := map[string]bool{
			"GET":    true,
			"POST":   true,
			"PUT":    true,
			"DELETE": true,
		}

		if !allowedMethods[c.Request.Method] {
			c.JSON(http.StatusMethodNotAllowed, gin.H{
				"error": "Method Not Allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
