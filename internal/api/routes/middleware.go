package routes

import (
	"Marketplace/internal/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Извлекаем токен из заголовков авторизации
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Проверяем, начинается ли заголовок с "Bearer "
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Декодируем токен
		token := tokenParts[1]
		secretKey := jwt.ProvideSecretKey()
		log.Printf("Secret Key: %s", string(secretKey))

		// Декодируем JWT токен
		decodedJWT, err := jwt.Decode(token, secretKey)
		if err != nil {
			log.Printf("Error decoding token: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Сохраняем данные пользователя в контексте запроса
		c.Set("user", decodedJWT)

		// Продолжаем выполнение запроса
		c.Next()
	}
}
