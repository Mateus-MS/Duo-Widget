package widget

import (
	"image"
	"image/color"
	"io"
	"log"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func InsertLabel(img *image.RGBA, streak string) {
	// Get the reusable font.Face
	fontFace := getFontFace()

	// Create a font.Drawer to draw text
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}), // White color
		Face: fontFace,
	}

	// Calculate text width
	textBounds, _ := font.BoundString(fontFace, streak)
	textWidth := (textBounds.Max.X - textBounds.Min.X).Ceil()

	// Calculate padding and positioning
	streakRightPadding := 8
	totalWidth := 47 + streakRightPadding + textWidth
	padding := (img.Bounds().Dx() - totalWidth) / 2

	// Set the position where the text will be drawn
	d.Dot = fixed.Point26_6{
		X: fixed.I(padding + 47 + streakRightPadding), // Horizontal position
		Y: fixed.I(81),                                // Vertical position (baseline)
	}

	// Draw the string on the image
	d.DrawString(streak)
}

func getFontFace() font.Face {
	// Fetch the font file from the provided URL
	resp, err := http.Get("https://duo-widget.vercel.app/fonts/MADE_Tommy_Soft_Bold_PERSONAL_USE.otf")
	if err != nil {
		log.Fatalf("failed to fetch font: %v", err)
	}
	defer resp.Body.Close()

	// Read the font bytes
	fontBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read font: %v", err)
	}

	// Parse the font
	fontParsed, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	// Create a font face with specific options
	fontFace, err := opentype.NewFace(fontParsed, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	return fontFace
}
