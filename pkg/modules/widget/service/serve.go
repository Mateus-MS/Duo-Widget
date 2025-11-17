package widget_service

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"strconv"

	_ "embed"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

//go:embed files/font.otf
var fontBytes []byte

//go:embed files/streak_logo.png
var streakLogoBytes []byte

// TODO: this works, but needs a good refactor

var bgColor = color.RGBA{R: 255, G: 35, B: 42, A: 0}
var streakTextColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
var imgSize = 300

func (serv *service) Serve(username, mood string) ([]byte, error) {
	streakInt, err := serv.externalService.GetStreak(username)
	if err != nil {
		return nil, err
	}
	streak := strconv.Itoa(streakInt)

	// Check in "temp" folder if already have a image with the same `mood` and `streak`.
	data, err := serv.repository.ReadFromCache(mood, streak)
	if err == nil {
		return data, nil
	}

	// 2.0 - If not, Create a new image, then serve it

	// Create the canvas
	canvas := image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{C: bgColor}, image.Point{}, draw.Src)

	// Insert the moodRaw
	moodRaw := serv.widgetRaw[mood]
	draw.Draw(canvas, moodRaw.Bounds(), moodRaw, image.Point{}, draw.Over)

	// Insert the streak
	streakLogo, err := getStreakLogo()
	insertStreak(canvas, streakLogo, streak)

	// Encode as jpg
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, canvas, nil)
	if err != nil {
		return nil, err
	}

	// Then save it in "temp" folder.
	err = serv.repository.SaveInCache(mood, streak, buf.Bytes())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// The streak is another image builded with the streak logo and the text with the streak count
func insertStreak(canvas *image.RGBA, streakLogo image.Image, streakCount string) {
	padding := 12
	streakCountWidth := measureTextWidth(streakCount)
	streakCountHeight := 57
	streakLogoWidth := streakLogo.Bounds().Dx()

	// Create the streak canvas
	// It's size is calculated to be enough to fit the logo and the label with a padding in between
	streakCanvas := image.NewRGBA(image.Rect(0, 0, streakLogoWidth+padding+streakCountWidth, streakCountHeight))

	// Insert the logo
	draw.Draw(streakCanvas, streakLogo.Bounds(), streakLogo, image.Point{0, 0}, draw.Over)

	// Insert the label
	d, err := getFontDrawer(streakCanvas, streakTextColor)
	if err != nil {
		return
	}

	d.Dot = fixed.Point26_6{fixed.Int26_6(streakLogoWidth+padding) << 6, fixed.Int26_6(streakCountHeight) << 6}
	d.DrawString(streakCount)

	// Insert the streak canvas into the original canvas
	draw.Draw(canvas, streakCanvas.Bounds().Add(image.Point{(canvas.Bounds().Dx() - streakCanvas.Bounds().Dx()) / 2, 20}), streakCanvas, image.Point{0, 0}, draw.Over)
}

func measureTextWidth(label string) int {
	d, err := getFontDrawer(nil, nil)
	if err != nil {
		return -1
	}

	textWidth := d.MeasureString(label).Round()
	return textWidth
}

func getFontDrawer(img *image.RGBA, color color.Color) (*font.Drawer, error) {
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		return nil, errors.New("failed to parse font")
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		return nil, errors.New("failed to create face")
	}

	return &font.Drawer{
		Dst:  img,
		Src:  &image.Uniform{C: color},
		Face: face,
	}, nil
}

func getStreakLogo() (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(streakLogoBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}
