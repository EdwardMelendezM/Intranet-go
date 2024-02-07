package main

import (
	"auth/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(corsMiddleware())

	router.POST("api/v1/login", loginHandler)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

// Función que crea el middleware CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func loginHandler(context *gin.Context) {
	var loginBody domain.LoginBody

	if err := context.ShouldBindJSON(&loginBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if loginBody.Username == "usuario" && loginBody.Password == "contrasena" {
		context.JSON(http.StatusOK, loginBody)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciales incorrectas"})
	}
}
