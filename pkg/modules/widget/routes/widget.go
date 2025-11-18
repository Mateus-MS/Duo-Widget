package widget_routes

import (
	"errors"
	"net/http"
	"strings"

	widget_service "github.com/Mateus-MS/Duo-Widget/modules/widget/service"
	"github.com/gin-gonic/gin"
)

func WidgetRoute(service widget_service.Iservice) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		if username == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("Must provide your username"))
			return
		}

		mood := c.Param("mood")
		if mood == "" {
			mood = "angry"
		}
		mood = strings.Replace(mood, ".png", "", 1)

		data, err := service.Serve(username, mood)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// Configure the response header to be cached by github
		c.Header("Cache-Control", "public, max-age=3600")
		c.Data(http.StatusOK, "image/png", data)
	}
}
