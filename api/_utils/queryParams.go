package utils

import (
	"net/http"
)

func QueryFromURL(param string, r *http.Request) (string, error) {
	return r.URL.Query().Get(param), nil
}
