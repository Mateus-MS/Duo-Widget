package widget_routes

import (
	"errors"
	"net/http"

	widget_service "github.com/Mateus-MS/Duo-Widget/modules/widget/service"
	"github.com/gin-gonic/gin"
)

func WidgetRoute(service widget_service.Iservice) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("Must provide your username"))
			return
		}

		mood := c.Query("mood")
		if mood == "" {
			mood = "angry"
		}

		data, err := service.Serve(username, mood)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Writer.Write(data)
	}
}
