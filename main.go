package main

import (
	"fmt"
	"log"
	"os"
	"weventure_test/api/controllers"
	"weventure_test/common/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db.ConnectDB()

	var (
		r               = gin.New()
		authController  = controllers.AuthController{}
		tasksController = controllers.TasksController{}
		usersController = controllers.UsersController{}
	)
	authController.Init(r)
	tasksController.Init(r)
	usersController.Init(r)
	r.Run(fmt.Sprintf(":%s", os.Getenv("API_PORT")))
}
