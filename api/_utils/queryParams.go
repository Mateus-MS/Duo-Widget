package utils

import (
	"errors"
	"net/http"
)

func QueryFromURL(param string, r *http.Request) (string, error) {
	str := r.URL.Query().Get(param)

	if str == "" {
		return "", errors.New("Parameter Not Found")
	}
	return str, nil

}
