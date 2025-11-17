package main

import (
	external_repository "github.com/Mateus-MS/Duo-Widget/modules/external/repository/inMemory"
	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	widgets_images "github.com/Mateus-MS/Duo-Widget/modules/widget/images"
	widget_repository "github.com/Mateus-MS/Duo-Widget/modules/widget/repository/local"
	widget_schedule "github.com/Mateus-MS/Duo-Widget/modules/widget/schedule"
	widget_service "github.com/Mateus-MS/Duo-Widget/modules/widget/service"
	"github.com/Mateus-MS/Duo-Widget/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	widgetRaw, err := widgets_images.New()
	if err != nil {
		panic(err)
	}

	externalService := external_service.New(external_repository.New())
	widgetService := widget_service.New(widget_repository.New(), *widgetRaw, externalService)

	widget_schedule.StartCleanCacheSchedule(widgetService)

	routes.Init(router, widgetService, externalService)
	router.Run(":9898")
}
