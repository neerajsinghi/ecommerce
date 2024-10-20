package habndlerfunction

import (
	"ecommerceuser/auth/controller"
	"ecommerceuser/auth/service"
	"ecommerceuser/model"
	"ecommerceuser/repository"

	"github.com/gin-gonic/gin"
)

var controllerU = controller.NewAuthController(*service.NewAuthService(*repository.NewUserRepository(model.GetDB()), model.GetJWTSecret(), model.GetJWTExpiration()))

func LoginHandler(c *gin.Context) {
	controllerU.Login(c)
}

func RegisterHandler(c *gin.Context) {
	controllerU.Register(c)
}

func LogoutHandler(c *gin.Context) {
	controllerU.Logout(c)
}
