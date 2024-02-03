package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"auth/entities"
)

func main() {
	router := gin.Default()

	router.POST("api/v1/login", loginHandler)

	err := router.Run(":8080")
	if err != nil {
		return
	}
}

func loginHandler(context *gin.Context) {
	var loginBody entities.LoginBody

	if err := context.ShouldBindJSON(&loginBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if loginBody.Username == "usuario" && loginBody.Password == "contrasena" {
		context.JSON(http.StatusOK, loginBody)
	} else {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Credenciales incorrectas"})
	}
}
