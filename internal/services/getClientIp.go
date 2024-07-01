package services

import (
	"errors"
	// "net"
	// "strings"

	"github.com/gin-gonic/gin"
)

func GetClientIp(c *gin.Context) (string, error){

	ipAddr := c.ClientIP()
	if len(ipAddr) != 0 && ipAddr != "::1" {
		return ipAddr, nil
	}
	return "", errors.New("could not get client IP address")
}