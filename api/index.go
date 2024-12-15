package handler

import (
	"fmt"
	"image/jpeg"
	"net/http"

	utils "github.com/Mateus-MS/Duo-Widget/_utils"
	gee "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.New()

	server.GET("/api", func(context *gee.Context) {

		mood, err := utils.QueryFromURL("mood", r)
		if err != nil {
			context.Writer.WriteHeader(http.StatusBadRequest)
			context.JSON(http.StatusBadRequest, gee.H{"error": "No mood passed"})
			return
		}

		imageName := utils.GetMood(mood)
		img, err := utils.QueryImage(fmt.Sprintf("d%s.jpg", imageName))
		if err != nil {
			context.Writer.WriteHeader(http.StatusBadRequest)
			context.JSON(http.StatusBadRequest, gee.H{"error": "Failed to fetch image"})
			return
		}

		context.Writer.Header().Set("Content-Type", "image/jpeg")

		err = jpeg.Encode(context.Writer, img, nil)
		if err != nil {
			context.Writer.WriteHeader(http.StatusInternalServerError)
			context.JSON(500, gee.H{"error": "Failed to encode image"})
			return
		}
	})

	server.Handle(w, r)
}
