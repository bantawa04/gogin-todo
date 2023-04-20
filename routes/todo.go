package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func TodoRoutes() *gin.Engine {
	r := gin.Default()

	// Set up routes
	r.GET("/todos", controllers.ListTodos)
	r.GET("/todo/:id", controllers.GetTodo)
	r.POST("/todo/create", controllers.CreateTodo)
	r.PUT("/todo/:id", controllers.UpdateTodo)
	r.DELETE("/todo/:id", controllers.DeleteTodo)
	return r
}
