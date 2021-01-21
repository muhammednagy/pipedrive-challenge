package pipedrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/muhammednagy/pipedirve-challenge/models"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

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

// Create a new person in Pipedrive which should have the same name as the github username
func CreatePerson(username, token string) int {
	var createPersonResponse models.PipedriveCreatePersonResponse
	parameters := map[string]string{
		"name": username,
	}
	parametersJSON, _ := json.Marshal(parameters)

	body, err := makeRequest(token, "persons", parametersJSON)
	if err != nil {
		log.Error("Failed to create person in Pipedrive error: ", err)
		return 0
	}
	err = json.Unmarshal(body, &createPersonResponse)
	if err != nil {
		log.Error("Failed to create person in Pipedrive error: ", err)
		return 0
	}
	if !createPersonResponse.Success {
		log.Error("Failed to create person in Pipedrive error: ", createPersonResponse.Error)
		return 0
	}
	return createPersonResponse.Data.ID
}

// Create a new activity connected to a person with subject and a note to hold gist info and its files links or
// pull url if it's truncated due to big size
func CreateActivity(subject, note, token string, personID int) error {
	var status models.PipedriveResponseStatus
	parameters := struct {
		Subject  string `json:"subject"`
		Note     string `json:"note"`
		PersonId int    `json:"person_id"`
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
	return fmt.Errorf(status.Error)
}
