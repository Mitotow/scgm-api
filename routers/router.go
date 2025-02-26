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
	apiV1group := r.Group("/api/v1")

	// Locations routes
	locationRepository := repositories.NewLocationsRepositoryImpl(db)
	locationsService := services.NewLocationsService(locationRepository)
	locationsGroup := apiV1group.Group("/locations")
	// TODO: Add group.use() for jwt middleware
	locationsGroup.GET("", func(c *gin.Context) { v1.GetLocations(c, locationsService) })
	locationsGroup.GET("/:name", func(c *gin.Context) { v1.GetLocationByName(c, locationsService) })

	// TODO: SET AS API ADMIN ROUTES
	locationsGroup.POST("", func(c *gin.Context) {})
	locationsGroup.PUT("", func(c *gin.Context) {})
	locationsGroup.DELETE("", func(c *gin.Context) {})

	// Missions routes
	missionsGroup := apiV1group.Group("/missions")
	missionsGroup.GET("", func(c *gin.Context) {})
	missionsGroup.GET("/:id", func(c *gin.Context) {})

	// TODO: SET AS GUILD ADMIN ROUTES
	missionsGroup.POST("", func(c *gin.Context) {})
	missionsGroup.PUT("", func(c *gin.Context) {})
	missionsGroup.DELETE("", func(c *gin.Context) {})

	// Guild routes
	guildGroup := apiV1group.Group("/guild")
	guildGroup.GET("", func(c *gin.Context) {})

	// TODO: SET AS GUILD OWNER ROUTE
	guildGroup.PUT("", func(c *gin.Context) {})

	// GuildStorage routes
	guildStorageGroup := guildGroup.Group("/storage")
	guildStorageGroup.GET("", func(c *gin.Context) {})

	// TODO: SET AS ??
	guildStorageGroup.POST("", func(c *gin.Context) {})
	guildStorageGroup.PUT("", func(c *gin.Context) {})
	guildStorageGroup.DELETE("", func(c *gin.Context) {})

	// GuildEconomy routes
	guildEconomyGroup := guildGroup.Group("/economy")
	guildEconomyGroup.GET("", func(c *gin.Context) {})

	// TODO: SET AS ??
	guildEconomyGroup.POST("", func(c *gin.Context) {})
	guildEconomyGroup.PUT("", func(c *gin.Context) {})
	guildEconomyGroup.DELETE("", func(c *gin.Context) {})

	return r
}
