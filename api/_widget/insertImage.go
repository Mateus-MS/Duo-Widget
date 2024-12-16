package widget

import (
	"image"
	"image/draw"
)

func InsertPNGImage(canvas *image.RGBA, img image.Image, x, y int) {
	srcBounds := img.Bounds()
	// Define the destination rectangle
	dstRect := image.Rect(x, y, x+srcBounds.Dx(), y+srcBounds.Dy())

	// Draw the image onto the canvas
	draw.Draw(canvas, dstRect, img, image.Point{}, draw.Over)
}

func InsertJPGImage(canvas *image.RGBA, img image.Image, x, y int) {
	srcBounds := img.Bounds()
	// Define the destination rectangle
	dstRect := image.Rect(x, y, x+srcBounds.Dx(), y+srcBounds.Dy())

	// Draw the image onto the canvas
	draw.Draw(canvas, dstRect, img, image.Point{}, draw.Src)
}
