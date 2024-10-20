package habndlerfunction

import (
	"ecommerceuser/model"
	"ecommerceuser/repository"
	"ecommerceuser/user/controller"
	"ecommerceuser/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var controllerU = controller.NewUserController(service.NewUserService(*repository.NewUserRepository(model.GetDB())))

func GetUser(c *gin.Context) {
	id := c.GetInt("userID")
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	controllerU.GetUserByID(c, id)
}

func GetUsers(c *gin.Context) {
	if c.GetString("role") != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	controllerU.GetAllUsers(c)
}

func UpdateUser(c *gin.Context) {
	id := c.GetInt("userID")
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uint(id)
	controllerU.UpdateUser(c, user)
}

func DeleteUser(c *gin.Context) {
	id := c.GetInt("userID")
	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	controllerU.DeleteUser(c, id)
}
