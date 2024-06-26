package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context){
	name := c.Query("visitor_name")
	location := ""
	clientIp := c.ClientIP()


	c.JSON(200, gin.H{
		"client_ip": clientIp,
		"location": location,
		"greeting": fmt.Sprintf("Hello, %s!", name),
	})
}