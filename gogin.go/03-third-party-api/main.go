package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var apiURL = "http://date.jsontest.com/"

func main() {
	fmt.Println("Third Party API DATA")
	r := gin.Default()
	r.GET("/get", getValue)
	r.Run(":9090")
}

func getValue(c *gin.Context) {
	resp, err := http.Get(apiURL)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	var target map[string]interface{}
	if err = json.Unmarshal(data, &target); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": target,
	})
}
