package handler

import (
	"image/png"
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
			// No mood passed
			// Continue with default
			mood = "happy"
		}

		// Get userID parameter
		userID, err := utils.QueryFromURL("userID", r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the image referenced to the mood
		img, _ := widget.CreateWidget(utils.GetMood(mood), utils.QueryStreak(userID))

		// Encode the image to send back
		w.Header().Set("Content-Type", "image/png")

		err = png.Encode(context.Writer, img)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	})

	server.Handle(w, r)
}
