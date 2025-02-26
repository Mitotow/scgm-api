package v1

import (
	"github.com/Mitotow/scgm-api/config"
	"github.com/Mitotow/scgm-api/repositories"
	"github.com/Mitotow/scgm-api/services"
	"github.com/gin-gonic/gin"
	"strconv"
)

var locationRepository = repositories.NewLocationsRepositoryImpl(config.DatabaseConnection())

// GetLocations
// @Summary Get a list of locations with pagination
// @Description Retrieve a paginated list of locations
// @Produce json
// @Success 200 {object} gin.H "Returns a JSON object containing the status, total number of locations, and the list of locations"
// @Failure 500 {object} gin.H "Returns a JSON object with an error message"
// @Router /api/v1/locations [get]
func GetLocations(c *gin.Context, service services.LocationsService) {
	var page int
	query, exists := c.GetQuery("page")
	queryAsInt, parseError := strconv.ParseInt(query, 10, 32)
	if !exists || parseError != nil || queryAsInt <= 0 {
		page = 1
	} else {
		page = int(queryAsInt)
	}

	response, err := service.FindAll(page)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(response.Status, response)
}

// GetLocationByName
// @Summary Get one location by name
// @Description Retrieve a specific location by its name
// @Produce json
// @Param name path string true "Name of the location"
// @Success 200 {object} gin.H "Returns a JSON object containing the status and the location details"
// @Failure 404 {object} gin.H "Returns a JSON object with an error message if the location is not found"
// @Router /api/v1/locations/{name} [get]
func GetLocationByName(c *gin.Context, service services.LocationsService) {
	name := c.Param("name")
	response, err := service.FindByName(name)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(response.Status, response)
}
