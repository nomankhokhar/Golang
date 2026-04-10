package routes

import (
	"todo-crud/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter initialises the Gin engine and registers all routes
func SetupRouter(tc *controllers.TodoController) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/todos", tc.GetAllTodos)
		api.GET("/todos/:id", tc.GetTodoByID)
		api.POST("/todos", tc.CreateTodo)
		api.PUT("/todos/:id", tc.UpdateTodo)
		api.DELETE("/todos/:id", tc.DeleteTodo)
	}

	return r
}
