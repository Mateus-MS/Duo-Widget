package external_repository

func (repo *repository) Reset() {
	repo.collection = make(map[string]int)
}
