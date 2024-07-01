package services

import (
	"errors"
	// "net"
	// "strings"

	"github.com/gin-gonic/gin"
)

func GetClientIp(c *gin.Context) (string, error){
	// xForwardedFor := strings.TrimSpace(c.Request.Header.Get("X-FORWARDED-FOR"))
	// remoteAddr := strings.TrimSpace(c.Request.Header.Get("REMOTE-ADDR"))
	// clientIP := strings.TrimSpace(c.Request.Header.Get("CLIENT-IP"))

	// ipAddr := ""
	// if len(xForwardedFor) != 0 {
	// 	ipAddr = xForwardedFor
	// } else if len(remoteAddr) != 0 {
	// 	ipAddr = remoteAddr
	// } else if len(clientIP) != 0 {
	// 	ipAddr = clientIP
	// }
	// if len(ipAddr) != 0 {
	// 	ip := net.ParseIP(ipAddr)
	// 	if ip == nil || ip.IsLoopback() {
	// 		return "", errors.New("invalid IP or Loopback IP address")
	// 	}
	// 	ip = ip.To4()
	// 	if ip == nil {
	// 		return "", errors.New("could not get IPv4 address")
	// 	}
	// 	return ipAddr, nil
	// }

	ipAddr := c.ClientIP()
	if len(ipAddr) != 0 && ipAddr != "::1" {
		return ipAddr, nil
	}
	return "", errors.New("could not get client IP address")
}