package handlers

import (
	"fmt"
	"hng11task1/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context){
	name := c.Query("visitor_name")
	if name == ""{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not get your name from the url query",
		})
		return
	}

	clientIP, err := services.GetClientIp(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// locDetails, err := services.GetLocationByIP(clientIP)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not get your location from IP",
	// 	})
	// 	return
	// }

	temp, err := services.GetTempByLoc(clientIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get your Temperature",
		})
		return
	}

	c.JSON(200, gin.H{
		"client_ip": clientIP,
		"location": temp["location"]["region"],
		"greeting": fmt.Sprintf("Hello, %s! the temperature is %.2f degrees Celcius in %s", name, temp["current"]["temp_c"], temp["location"]["region"]),
	})
}