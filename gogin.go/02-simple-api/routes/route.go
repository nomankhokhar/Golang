package routes

import (
	"simple-api/controller"

	"github.com/gin-gonic/gin"
)

func ApiRoute() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api/v1")
	api.GET("/get", controller.GetValues)
	api.POST("/post", controller.PostValues)
	api.PUT("/put", controller.PutValues)
	api.DELETE("/delete", controller.DeleteValues)

	return r
}
