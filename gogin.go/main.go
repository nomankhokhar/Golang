// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type Todo struct {
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Status string `json:"status"`
// }

// var todos = []Todo{
// 	{ID: 1, Title: "Learn Golang", Status: "Incomplete"},
// 	{ID: 2, Title: "Build API", Status: "Incomplete"},
// }

// func main() {
// 	r := gin.Default()

// 	// List all todos
// 	r.GET("/todos", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, todos)
// 	})

// 	// Get a single todo
// 	r.GET("/todos/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		for _, todo := range todos {
// 			if fmt.Sprintf("%d", todo.ID) == id {
// 				c.JSON(http.StatusOK, todo)
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 	})

// 	// Create a todo
// 	r.POST("/todos", func(c *gin.Context) {
// 		var newTodo Todo
// 		if err := c.ShouldBindJSON(&newTodo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		newTodo.ID = len(todos) + 1
// 		todos = append(todos, newTodo)
// 		c.JSON(http.StatusCreated, newTodo)
// 	})

// 	// Update a todo
// 	r.PUT("/todos/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		var updatedTodo Todo
// 		if err := c.ShouldBindJSON(&updatedTodo); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		for i, todo := range todos {
// 			if fmt.Sprintf("%d", todo.ID) == id {
// 				todos[i] = updatedTodo
// 				c.JSON(http.StatusOK, gin.H{"message": "Todo updated"})
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 	})

// 	// Delete a todo
// 	r.DELETE("/todos/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		for i, todo := range todos {
// 			if fmt.Sprintf("%d", todo.ID) == id {
// 				todos = append(todos[:i], todos[i+1:]...)
// 				c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
// 				return
// 			}
// 		}
// 		c.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
// 	})

// 	r.Run(":8080")
// }

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type api struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// var data api

// func main() {
// 	fmt.Println("Welcome to Golang->API")
// 	r := gin.Default()
// 	r.GET("/get", getValues)
// 	r.POST("/post", postValues)
// 	r.PUT("/put", putValues)
// 	r.DELETE("/delete", deleteValues)
// 	r.Run(":3000")
// }

// func getValues(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": data,
// 	})
// }

// func postValues(c *gin.Context) {
// 	err := c.BindJSON(&data)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Something Wrong",
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": data,
// 	})
// }

// func putValues(c *gin.Context) {
// 	err := c.BindJSON(&data)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{
// 			"message": "Something Wrong",
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": data,
// 	})
// }

// func deleteValues(c *gin.Context) {
// 	data = api{}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": data,
// 	})
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	fmt.Println("Third Party API DATA")
// 	r := gin.Default()
// 	r.GET("/get", getValue)
// 	r.Run(":9090")
// }

// var Url = "http://date.jsontest.com/"

// func getValue(c *gin.Context) {
// 	resp, err := http.Get(Url)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"message": err,
// 		})
// 	}

// 	defer resp.Body.Close()

// 	data, _ := ioutil.ReadAll(resp.Body)

// 	var target map[string]interface{}

// 	err = json.Unmarshal(data, &target)
// 	if err != nil {
// 		c.JSON(400, gin.H{
// 			"Message": err,
// 		})
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": target,
// 	})
// }

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// )

// type Todo struct {
// 	ID    int    `json:"id"`
// 	Title string `json:"title"`
// 	Done  bool   `json:"done"`
// 	Body  string `json:"body"`
// }

// func main() {
// 	fmt.Print("Hello world")

// 	app := fiber.New()

// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "http://localhost:3000",
// 		AllowHeaders: "Origin, Content-Type, Accept",
// 	}))

// 	todos := []Todo{}

// 	app.Get("/healthcheck", func(c *fiber.Ctx) error {
// 		return c.SendString("OK")
// 	})

// 	app.Post("/api/todos", func(c *fiber.Ctx) error {
// 		todo := &Todo{}

// 		if err := c.BodyParser(todo); err != nil {
// 			return err
// 		}

// 		todo.ID = len(todos) + 1

// 		todos = append(todos, *todo)

// 		return c.JSON(todos)

// 	})

// 	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
// 		id, err := c.ParamsInt("id")

// 		if err != nil {
// 			return c.Status(401).SendString("Invalid id")
// 		}

// 		for i, t := range todos {
// 			if t.ID == id {
// 				todos[i].Done = true
// 				break
// 			}
// 		}

// 		return c.JSON(todos)
// 	})

// 	app.Get("/api/todos", func(c *fiber.Ctx) error {
// 		return c.JSON(todos)
// 	})

// 	log.Fatal(app.Listen(":4000"))

// }

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type api struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var data api

func main() {
	fmt.Println("Welcome to Golang->API")
	r := gin.Default()
	r.GET("/get", getValues)
	r.POST("/post", postValues)
	r.PUT("/put", putValues)
	r.DELETE("/delete", deleteValues)
	r.Run(":3000")
}

func getValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func postValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func putValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func deleteValues(c *gin.Context) {
	data = api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
