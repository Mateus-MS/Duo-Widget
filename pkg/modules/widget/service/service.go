package widget_service

import (
	external_service "github.com/Mateus-MS/Duo-Widget/modules/external/service"
	widgets_images "github.com/Mateus-MS/Duo-Widget/modules/widget/images"
	widget_repository "github.com/Mateus-MS/Duo-Widget/modules/widget/repository/local"
)

type Iservice interface {
	Serve(string, string) ([]byte, error)
}
type service struct {
	externalService external_service.Iservice
	repository      widget_repository.Irepository
	widgetRaw       widgets_images.MoodRaw
}

func New(repo widget_repository.Irepository, widgetRaw widgets_images.MoodRaw, externalService external_service.Iservice) *service {
	return &service{
		externalService: externalService,
		repository:      repo,
		widgetRaw:       widgetRaw,
	}
}
