package controllers

import (
	"net/http"

	"weventure_test/api/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type AuthController struct {
}

func (a *AuthController) Init(r *gin.Engine) {
	var (
		authService = services.AuthService{}
	)
	r.POST("/login", func(c *gin.Context) {
		var body authBody
		c.ShouldBindBodyWith(&body, binding.JSON)
		var res = authService.Login(body.ID, body.Password)
		var status int
		if res.Status == 1 {
			status = http.StatusOK
		} else {
			status = http.StatusUnauthorized
		}
		c.JSON(status, res)
	})
}

type authBody struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
