package main

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "arna/docs"

	"arna/internal/handlers"
	"arna/internal/store"
)

// @title Task API
// @version 1.0
// @description Minimal task management API
// @host localhost:8080
// @BasePath /
func main() {

	r := gin.Default()

	store := store.NewMemoryStore()
	handler := handlers.NewTaskHandler(store)

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.ListTasks)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}