package v1

import (
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

var locationRepository = repositories.NewLocationsRepositoryImpl(config.DatabaseConnection())

// GetLocations
// @Summary Get a list of locations with pagination
// @Produce json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/locations
func GetLocations(c *gin.Context) {
	locations := locationRepository.FindAll()
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"total":   len(locations),
		"results": locations,
	})
}
