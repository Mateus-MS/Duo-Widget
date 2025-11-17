package widget_repository

import (
	"fmt"
	"os"
)

func (repo *repository) ReadFromCache(mood, streak string) ([]byte, error) {
	data, err := os.ReadFile(fmt.Sprintf("./cache/%s/%s.jpg", mood, streak))
	if err != nil {
		return nil, err
	}

	return data, nil
}
