package user

import (
	"ecommerceuser/middleware"
	habndlerfunction "ecommerceuser/user/habndlerFunction"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	authorized := r.Group("/user")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/id", habndlerfunction.GetUser)
		authorized.GET("/", habndlerfunction.GetUsers)
		authorized.PATCH("/", habndlerfunction.UpdateUser)
		authorized.DELETE("/", habndlerfunction.DeleteUser)
	}
}
