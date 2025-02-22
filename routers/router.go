package routers

import "github.com/gin-gonic/gin"

func CreateRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	locationsGroup := r.Group("/api/v1/locations")

	// TODO: Add group.use() for jwt middleware

	locationsGroup.GET("/")

	return r
}
