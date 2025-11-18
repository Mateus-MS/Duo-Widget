package widgets_images

import (
	"bytes"
	"embed"
	"image"
	"math/rand/v2"
	"regexp"

	_ "image/jpeg"
	_ "image/png"
)

//go:embed files/*
var widgetFiles embed.FS

type MoodRaw map[string]image.Image

// New loads all embedded images into a map
func New() (*MoodRaw, error) {
	m := make(MoodRaw)

	// List all embedded image files
	entries, err := widgetFiles.ReadDir("files")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		fileName := entry.Name()

		img, err := getImage(fileName)
		if err != nil {
			return nil, err
		}

		// Remove the extension for the map key
		re := regexp.MustCompile(`\..*`)
		name := re.ReplaceAllString(fileName, "")

		m[name] = img
	}

	return &m, nil
}

// getImage reads and decodes an image from the embedded FS
func getImage(fileName string) (image.Image, error) {
	data, err := widgetFiles.ReadFile("files/" + fileName)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}

// TODO: make that it doesn't build the array with the keys on every call
func (mr MoodRaw) GetRandom() string {
	if len(mr) == 0 {
		return ""
	}

	keys := make([]string, 0, len(mr))
	for mood := range mr {
		keys = append(keys, mood)
	}

	return keys[rand.IntN(len(keys))]
}
