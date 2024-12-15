package utils

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func QueryStreak(userID string) string {

	response, err := http.Get(fmt.Sprintf("https://www.duolingo.com/2017-06-30/users/%s", userID))
	if err != nil {
		return ""
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	regex := regexp.MustCompile(`"streak":\s*(\d+)`)
	match := regex.FindStringSubmatch(string(body))
	if len(match) == 0 {
		return ""
	}
	return match[1]

}
