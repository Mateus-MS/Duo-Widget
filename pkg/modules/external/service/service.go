package external_service

type Iservice interface {
	GetStreak(string) (int, error)
}
type service struct {
}

func New() *service {
	return &service{}
}
