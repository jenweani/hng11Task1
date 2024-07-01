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

	locDetails, err := services.GetLocationByIP(clientIP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get your location from IP",
		})
		return
	}

	temp, err := services.GetTempByLoc(locDetails.Latitude, locDetails.Longitude)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not get your Temperature",
		})
		return
	}

	c.JSON(200, gin.H{
		"client_ip": clientIP,
		"location": locDetails.City,
		"greeting": fmt.Sprintf("Hello, %s! the temperature is %s degrees Celcius in %s", name, temp, locDetails.City),
	})
}