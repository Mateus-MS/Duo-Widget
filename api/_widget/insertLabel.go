package widget

import (
	"errors"
	"image"
	"image/color"
	"io"
	"net/http"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func InsertLabel(img *image.RGBA, x, y int, label string, col color.Color) {
	// Font face
	d, err := getFontDrawer(img, x, y, col)
	if err != nil {
		return
	}

	// Draw
	d.Dot = fixed.Point26_6{fixed.Int26_6(x) << 6, fixed.Int26_6(y) << 6}
	d.DrawString(label)
}

func MeasureText(img *image.RGBA, x, y int, label string, col color.Color) int {
	d, err := getFontDrawer(img, x, y, col)
	if err != nil {
		return -1
	}

	textWidth := d.MeasureString(label).Round()

	return textWidth
}

func getFontDrawer(img *image.RGBA, x, y int, col color.Color) (*font.Drawer, error) {
	d := &font.Drawer{}

	resp, err := http.Get("https://duo-widget.vercel.app/fonts/MADE_Tommy_Soft_Bold_PERSONAL_USE.otf")
	if err != nil {
		return d, err
	}

	fontBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return d, errors.New("failed to parse font")
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return d, errors.New("failed to create new face")
	}

	d.Dst = img
	d.Src = &image.Uniform{C: col}
	d.Face = face

	return d, nil
}
