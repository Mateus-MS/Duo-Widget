package external_repository

type Irepository interface {
	SaveInCache(string, int) error
	ReadFromCache(string) (int, error)
	Reset()
}

type repository struct {
	collection map[string]int
}

func New() *repository {
	col := make(map[string]int)
	return &repository{
		collection: col,
	}
}
