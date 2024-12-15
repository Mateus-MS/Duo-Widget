package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"
)

func QueryImage(imageName string) (image.Image, error) {

	// Create the query URL
	imageURL := fmt.Sprintf("https://duo-widget.vercel.app/images/%s", imageName)
	var imageBits image.Image

	// Request the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return imageBits, err
	}
	defer resp.Body.Close()

	// Decode the image
	if strings.Contains(imageName, ".png") {
		imageBits, err = png.Decode(resp.Body)
	} else {
		imageBits, err = jpeg.Decode(resp.Body)
	}
	if err != nil {
		return imageBits, err
	}

	// Complete
	return imageBits, nil

}
