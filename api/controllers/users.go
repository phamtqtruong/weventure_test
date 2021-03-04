package controllers

import (
	"net/http"

	services "weventure_test/api/services"

	"weventure_test/common/rest"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
}

func (u *UsersController) Init(r *gin.Engine) {
	var (
		user        = r.Group("user")
		userService = services.UsersService{}
	)
	user.GET("/all", func(c *gin.Context) {
		var res = rest.Response{
			Code:   http.StatusOK,
			Status: 1,
			Data:   userService.GetAll(),
		}
		c.JSON(200, res)
	})
}
