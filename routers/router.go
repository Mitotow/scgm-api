package routers

import (
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/repositories"
	v1 "github.com/Mitotow/scgm-api/routers/api/v1"
	"github.com/Mitotow/scgm-api/services"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	db := config.DatabaseConnection()

	// Locations routes
	locationRepository := repositories.NewLocationsRepositoryImpl(db)
	locationsService := services.NewLocationsService(locationRepository)
	locationsGroup := r.Group("/api/v1/locations")
	// TODO: Add group.use() for jwt middleware
	locationsGroup.GET("", func(c *gin.Context) { v1.GetLocations(c, locationsService) })
	locationsGroup.GET("/:name", func(c *gin.Context) { v1.GetLocationByName(c, locationsService) })

	return r
}
