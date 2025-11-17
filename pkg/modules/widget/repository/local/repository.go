package widget_repository

type Irepository interface {
	SaveInCache(string, string, []byte) error
	ReadFromCache(string, string) ([]byte, error)
}

type repository struct {
}

func New() *repository {
	return &repository{}
}
