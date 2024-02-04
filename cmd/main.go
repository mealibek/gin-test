package main

import (
	"os"

	"github.com/gin-gonic/gin"
	handlers "github.com/mealibek/gin-test/handlers"
	"github.com/mealibek/gin-test/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/hello", handlers.Hello)

	r.POST("/api/v1/auth/signup", handlers.SignUp)
	r.POST("/api/v1/auth/signin", handlers.SignIn)

	err := r.Run(os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}
}
