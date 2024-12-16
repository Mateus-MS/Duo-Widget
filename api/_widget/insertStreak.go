package widget

import (
	"image"
	"image/color"

	utils "github.com/Mateus-MS/Duo-Widget/_utils"
)

func InsertStreak(img *image.RGBA, streak string) {

	// Get the streak logo
	streakLogo, err := utils.QueryImage("s.png")
	if err != nil {
		return
	}

	// Calculate positioning variables
	streakTextColor := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	streakTextWidth := MeasureText(img, 0, 0, streak, streakTextColor)
	streakRightPadding := 8

	totalWidth := 47 + streakRightPadding + streakTextWidth
	sidePaddings := (img.Bounds().Dx() - totalWidth) / 2

	// Draw the streak logo
	InsertPNGImage(img, streakLogo, sidePaddings, 27)

	// Draw the streak counter
	InsertLabel(img, sidePaddings+47+streakRightPadding, 81, streak, streakTextColor)

}
