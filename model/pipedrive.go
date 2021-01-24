package model

type PipedriveCreatePersonResponse struct {
	Data struct {
		ID int `json:"id"`
	} `json:"data"`
	PipedriveResponseStatus
}

type PipedriveResponseStatus struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}
