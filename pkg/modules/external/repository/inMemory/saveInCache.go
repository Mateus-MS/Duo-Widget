package external_repository

func (repo *repository) SaveInCache(username string, streak int) error {
	repo.collection[username] = streak
	return nil
}
