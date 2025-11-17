package external_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	external_model "github.com/Mateus-MS/Duo-Widget/modules/external/model"
)

func (serv *service) GetStreak(username string) (int, error) {
	// Create the request
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://www.duolingo.com/2017-06-30/users?username=%s", username), nil)
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Decode the response
	var respData = external_model.RequestResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return 0, err
	}

	return respData.Users[0].Streak, nil
}
