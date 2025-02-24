package routers

import (
	v1 "github.com/Mitotow/scgm-api/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	locationsGroup := r.Group("/api/v1/locations")

	// TODO: Add group.use() for jwt middleware

	locationsGroup.GET("", v1.GetLocations)
	locationsGroup.GET("/:name", v1.GetLocationByName)

	return r
}
