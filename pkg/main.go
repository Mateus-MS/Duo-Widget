package main

import (
	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	widgets_images "github.com/Mateus-MS/Duo-Widget/modules/widget/images"
	widget_repository "github.com/Mateus-MS/Duo-Widget/modules/widget/repository/local"
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

	externalService := external_service.New()
	widgetService := widget_service.New(widget_repository.New(), *widgetRaw, externalService)

	routes.Init(router, widgetService, externalService)
	router.Run(":9898")
}
