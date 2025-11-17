package external_routes

import (
	"errors"
	"net/http"
	"strconv"

	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	"github.com/gin-gonic/gin"
)

// This is a utility route that returns only the `streak` data from official DuoLingo API
func StreakRoute(service external_service.Iservice) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if username == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("must provide your username"))
			return
		}

		data, err := service.GetStreak(username)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.String(http.StatusOK, strconv.Itoa(data))
	}
}
