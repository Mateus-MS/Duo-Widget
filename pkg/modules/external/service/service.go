package external_service

import external_repository "github.com/Mateus-MS/Duo-Widget/modules/external/repository/inMemory"

type Iservice interface {
	GetStreak(string) (int, error)
}
type service struct {
	repository external_repository.Irepository
}

func New(externalRepository external_repository.Irepository) *service {
	return &service{
		repository: externalRepository,
	}
}
