package widget

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"

	utils "github.com/Mateus-MS/Duo-Widget/_utils"
)

func CreateWidget(moodReference string, streak string) (*image.RGBA, error) {
	widgetImage := image.NewRGBA(image.Rect(0, 0, 300, 300))
	bgColor := color.RGBA{R: 255, G: 35, B: 42, A: 0}
	draw.Draw(widgetImage, widgetImage.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	moodImage, err := utils.QueryImage(fmt.Sprintf("d%s.jpg", moodReference))
	if err != nil {
		return widgetImage, errors.New("failed to fetch the image")
	}

	InsertJPGImage(widgetImage, moodImage, 0, 0)
	InsertStreak(widgetImage, streak)

	return widgetImage, nil
}
