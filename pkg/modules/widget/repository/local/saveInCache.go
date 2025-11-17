package widget_repository

import (
	"fmt"
	"os"
)

func (repo *repository) SaveInCache(mood string, streak string, image []byte) error {
	// Ensure the folder exist
	if err := os.MkdirAll(fmt.Sprintf("./modules/widget/cache/%s", mood), 0755); err != nil {
		return err
	}

	// Save the image
	return os.WriteFile(fmt.Sprintf("./modules/widget/cache/%s/%s.jpg", mood, streak), image, 0644)
}
