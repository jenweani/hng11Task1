package api

import (
	"os"
	"hng11task1/internal/handlers"

	"github.com/gin-contrib/cors"
	geo "github.com/cjgiridhar/gin-geo"
	"github.com/gin-gonic/gin"
)

func BuildRoutesHandler() *gin.Engine {
	r := gin.New()

	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.Use(geo.Default("github.com/cjgiridhar/gin-geo/db/GeoLite2-City.mmdb"))

	r.GET("/health", handlers.HealthHandler)
	r.GET("/api/hello", handlers.HelloHandler)

	return r
}