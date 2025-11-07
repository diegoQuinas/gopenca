package main

import (
	"github.com/gin-gonic/gin"
	"github.com/diegoQuinas/gopenca/config"
	"github.com/diegoQuinas/gopenca/handlers"
	"github.com/diegoQuinas/gopenca/repository"
	"github.com/diegoQuinas/gopenca/services"
)


func main() {
	db := config.ConnectDB()
	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userRepo, userService)

	r := gin.Default()
	r.GET("/users", userHandler.GetAll)
	r.POST("/users", userHandler.Create)
	r.PUT("/users/:id", userHandler.Update)
	r.DELETE("/users/:id", userHandler.Delete)
	

	r.Run(":8000")
}
