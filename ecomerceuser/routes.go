package main

import (
	"ecommerceuser/auth"
	"ecommerceuser/user"

	"github.com/gin-gonic/gin"
)

func addRoutes(apiGroup *gin.RouterGroup) {
	auth.RegisterRoutes(apiGroup)
	user.Routes(apiGroup)
}
