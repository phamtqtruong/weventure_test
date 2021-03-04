package controllers

import (
	services "weventure_test/api/services"

	"weventure_test/api/dto"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type TasksController struct {
}

func (t *TasksController) Init(r *gin.Engine) {
	var (
		user         = r.Group("task")
		tasksService = services.TasksService{}
	)
	user.GET("/list", func(c *gin.Context) {
		var res = tasksService.GetAll(c.Copy().Query("assignee"), c.Copy().Query("due_date"))
		c.JSON(200, res)
	})

	user.POST("/create", func(c *gin.Context) {
		var body dto.TaskDTO
		c.ShouldBindBodyWith(&body, binding.JSON)
		var res = tasksService.New(body)
		c.JSON(200, res)
	})
}
