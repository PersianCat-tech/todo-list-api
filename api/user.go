package api

import (
	"net/http"
	"todo_list/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	res := userRegister.Register()
	c.JSON(http.StatusAccepted, res)

}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	res := userLogin.Login()
	c.JSON(http.StatusAccepted, res)
}
