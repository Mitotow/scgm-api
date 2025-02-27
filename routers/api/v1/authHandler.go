package v1

import (
	"github.com/Mitotow/scgm-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context, authService services.AuthService) {
	redirectUrl := authService.Login()
	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

func Callback(c *gin.Context, authService services.AuthService) {
	response, err := authService.Callback(c.Query("code"))
	if err != nil {
		c.JSON(err.Status, err)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
