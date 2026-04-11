package controller

import (
	"net/http"
	"simple-api/model"

	"github.com/gin-gonic/gin"
)

var data model.Api

func GetValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func PostValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func PutValues(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something Wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}

func DeleteValues(c *gin.Context) {
	data = model.Api{}
	c.JSON(http.StatusOK, gin.H{
		"message": data,
	})
}
