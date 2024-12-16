package widget

import (
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func InsertLabel(img *image.RGBA, x, y int, label string, col color.Color) {
	// Font face
	d := getFontDrawer(img, x, y, col)

	// Draw
	d.DrawString(label)
}

func MeasureText(img *image.RGBA, x, y int, label string, col color.Color) int {
	d := getFontDrawer(img, x, y, col)

	textWidth := d.MeasureString(label).Round()

	return textWidth
}

func getFontDrawer(img *image.RGBA, x, y int, col color.Color) *font.Drawer {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	// Drawer
	d := &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: col},
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	return d
}
