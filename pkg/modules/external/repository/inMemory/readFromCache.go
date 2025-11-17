package external_repository

import "errors"

func (repo *repository) ReadFromCache(username string) (int, error) {
	result := repo.collection[username]
	if result == 0 {
		return 0, errors.New("user data is not in cache")
	}

	return result, nil
}
