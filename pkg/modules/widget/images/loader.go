package widgets_images

import (
	"image"
	"os"
	"regexp"

	_ "image/jpeg"
	_ "image/png"
)

type MoodRaw map[string]image.Image

var staticFolderPath string = "./static/imgs/widgets"

// Approximate map size: 0.45 MB
func New() (*MoodRaw, error) {
	m := make(MoodRaw)

	// Get all widgets base images paths
	entries, err := os.ReadDir(staticFolderPath)
	if err != nil {
		return nil, err
	}

	// On each widget base image
	for _, entry := range entries {
		fileName := entry.Name()

		img, err := getImage(fileName)
		if err != nil {
			return nil, err
		}

		// Remove the extension
		re := regexp.MustCompile(`\..*`)
		name := re.ReplaceAllString(fileName, "")

		m[name] = img
	}

	// Add the streak logo
	// img, err := getImage("streak_logo.png")
	// if err != nil {
	// 	return nil, err
	// }

	// m["streak_logo"] = img

	return &m, nil
}

func getImage(fileName string) (image.Image, error) {
	data, err := os.Open(staticFolderPath + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	img, _, err := image.Decode(data)
	if err != nil {
		return nil, err
	}

	return img, nil
}
