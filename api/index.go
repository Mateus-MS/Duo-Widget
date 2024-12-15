package handler

import (
	"io"
	"net/http"

	gee "github.com/tbxark/g4vercel"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gee.New()

	server.GET("/api", func(context *gee.Context) {
		imageURL := "https://duo-widget.vercel.app/images/test.jpg"

		resp, err := http.Get(imageURL)
		if err != nil {
			context.JSON(500, gee.H{"error": "Failed to fetch image"})
			return
		}
		defer resp.Body.Close()

		context.Writer.Header().Set("Content-Type", resp.Header.Get("Content-Type"))

		io.Copy(context.Writer, resp.Body)
	})

	server.Handle(w, r)
}
