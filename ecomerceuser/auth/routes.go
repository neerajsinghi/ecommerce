package auth

import (
	habndlerfunction "ecommerceuser/auth/habndlerFunction"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/login", habndlerfunction.LoginHandler)
	router.POST("/register", habndlerfunction.RegisterHandler)
	router.POST("/logout", habndlerfunction.LogoutHandler)
}
