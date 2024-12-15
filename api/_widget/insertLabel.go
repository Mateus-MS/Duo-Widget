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
	fontFace := getFontFace(img)

	textWidth := fontFace.MeasureString(streak).Round()

	streakRightPadding := 8

	totalWidth := 47 + streakRightPadding + textWidth
	padding := (img.Bounds().Dx() - totalWidth) / 2

	fontFace.Dot = fixed.Point26_6{fixed.Int26_6(padding+47+streakRightPadding) << 6, fixed.Int26_6(81) << 6}
	fontFace.DrawString(streak)
}

func getFontFace(img *image.RGBA) *font.Drawer {
	resp, err := http.Get("https://duo-widget.vercel.app/fonts/MADE_Tommy_Soft_Bold_PERSONAL_USE.otf")
	if err != nil {
		log.Fatal(err)
	}

	fontBytes, _ := io.ReadAll(resp.Body)

	f, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		log.Fatalf("failed to create new face: %v", err)
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
	}

	return d
}
