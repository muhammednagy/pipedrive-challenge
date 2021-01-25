package pipedrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CreatePersonResponse struct {
	Data struct {
		ID int `json:"id"`
	} `json:"data"`
	ResponseStatus
}

type ResponseStatus struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

// CreatePerson Create a new person in Pipedrive which should have the same name as the github username
func CreatePerson(username, token string) (int, error) {
	var createPersonResponse CreatePersonResponse
	parameters := map[string]string{
		"name": username,
	}
	parametersJSON, _ := json.Marshal(parameters)

	body, err := makeRequest(token, "persons", parametersJSON)
	if err != nil {
		return 0, fmt.Errorf("failed to create person in Pipedrive error: %s", err)
	}
	err = json.Unmarshal(body, &createPersonResponse)
	if err != nil {
		return 0, fmt.Errorf("failed to create person in Pipedrive error: %s", err)
	}
	if !createPersonResponse.Success {
		return 0, fmt.Errorf("failed to create person in Pipedrive error: %s", err)
	}
	return createPersonResponse.Data.ID, nil
}

// CreateActivity Create a new activity connected to a person with subject and a note to hold gist info and its files links or
// pull url if it's truncated due to big size
func CreateActivity(subject, note, token string, personID uint) error {
	var status ResponseStatus
	parameters := struct {
		Subject  string `json:"subject"`
		Note     string `json:"note"`
		PersonID uint   `json:"person_id"`
	}{
		subject,
		note,
		personID,
	}
	parametersJSON, _ := json.Marshal(parameters)

	body, err := makeRequest(token, "activities", parametersJSON)
	if err != nil {
		return fmt.Errorf("failed to create activity in Pipedrive error: %s", err)
	}
	err = json.Unmarshal(body, &status)
	if err != nil {
		return fmt.Errorf("failed to create activity in Pipedrive error: %s", err)
	}
	if !status.Success {
		return fmt.Errorf("failed to create activity in Pipedrive error: %s", status.Error)
	}
	return nil
}

const pipedriveBase = "https://api.pipedrive.com/v1/"

func makeRequest(token, path string, body []byte) ([]byte, error) {
	res, err := http.Post(pipedriveBase+path+"?api_token="+token, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to send request error: %s", err)
	}
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request response error: %s", err)
	}
	return body, nil
}
