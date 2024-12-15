package handler

import (
	"image/jpeg"
	"net/http"

	"github.com/Mateus-MS/Duo-Widget/utils"
	gee "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.New()

	server.GET("/api", func(context *gee.Context) {
		img, err := utils.QueryImage("test.jpg")
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
