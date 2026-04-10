package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: 1, Title: "Learn Golang", Status: "Incomplete"},
	{ID: 2, Title: "Build API", Status: "Incomplete"},
}

func main() {
	r := gin.Default()

	// List all todos
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, todos)
	})

	// Get a single todo
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				c.JSON(http.StatusOK, todo)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	// Create a todo
	r.POST("/todos", func(c *gin.Context) {
		var newTodo Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newTodo.ID = len(todos) + 1
		todos = append(todos, newTodo)
		c.JSON(http.StatusCreated, newTodo)
	})

	// Update a todo
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updatedTodo Todo
		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos[i] = updatedTodo
				c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	// Delete a todo
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	})

	r.Run(":8080")
}
