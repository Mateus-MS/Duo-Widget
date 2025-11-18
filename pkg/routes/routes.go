package routes

import (
	"net/http"

	external_routes "github.com/Mateus-MS/Duo-Widget/modules/external/routes"
	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	widget_routes "github.com/Mateus-MS/Duo-Widget/modules/widget/routes"
	widget_service "github.com/Mateus-MS/Duo-Widget/modules/widget/service"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine, widgetService widget_service.Iservice, externalService external_service.Iservice) {
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Alive")
	})

	router.GET("/:username/:mood", widget_routes.WidgetRoute(widgetService))

	router.GET("/streak/:username", external_routes.StreakRoute(externalService))
}
