package external_model

type RequestResponse struct {
	Users []struct {
		Streak int `json:"streak"`
	} `json:"users"`
}
