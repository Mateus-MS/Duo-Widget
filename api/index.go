package handler

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"net/http"

	utils "github.com/Mateus-MS/Duo-Widget/_utils"
	widget "github.com/Mateus-MS/Duo-Widget/_widget"
	gee "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.New()

	server.GET("/api", func(context *gee.Context) {

		// Get mood parameter
		mood, err := utils.QueryFromURL("mood", r)
		if err != nil {
			context.Writer.WriteHeader(http.StatusBadRequest)
			context.JSON(http.StatusBadRequest, gee.H{"error": "No mood passed"})
			return
		}

		// Get the image referenced to the mood
		moodReference := utils.GetMood(mood)
		img, err := utils.QueryImage(fmt.Sprintf("d%s.jpg", moodReference))
		if err != nil {
			context.Writer.WriteHeader(http.StatusBadRequest)
			context.JSON(http.StatusBadRequest, gee.H{"error": "Failed to fetch image"})
			return
		}

		// Add the days streak text into the image
		userID, err := utils.QueryFromURL("userID", r)
		if err != nil {
			context.Writer.WriteHeader(http.StatusBadRequest)
			context.JSON(http.StatusBadRequest, gee.H{"error": "Failed to get streak"})
			return
		}
		widget.InsertLabel(&ConvertToRGBA(img), utils.QueryStreak(userID))

		// Encode the image to send back
		err = jpeg.Encode(context.Writer, img, nil)
		if err != nil {
			context.Writer.WriteHeader(http.StatusInternalServerError)
			context.JSON(500, gee.H{"error": "Failed to encode image"})
			return
		}

	})

	server.Handle(w, r)
}

func ConvertToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)
	return rgba
}
