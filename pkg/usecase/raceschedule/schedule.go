package raceschedule

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Salaton/formula-one/pkg/domain"
)

const (
	responseTypeJSON        = "json"
	raceScheduleAPIEndpoint = "https://ergast.com/api/f1/%v.%v"
)

type RaceSchedule interface {
	GetSeasonRaceSchedules(ctx context.Context, year int) (*domain.DataResponse, error)
}

type ScheduleDetails struct{}

func NewRaceScheduleImplementation() *ScheduleDetails {
	return &ScheduleDetails{}
}

func (s ScheduleDetails) GetSeasonRaceSchedules(ctx context.Context, year int) (*domain.DataResponse, error) {
	var data *domain.DataResponse
	urlPath := fmt.Sprintf(raceScheduleAPIEndpoint, year, responseTypeJSON)
	resp, err := MakeRequest(ctx, http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make the api call %w", err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read request body: %w", err)
	}

	err = json.Unmarshal(responseData, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall data: %w", err)
	}

	return data, nil
}

// TODO: Extract this to a different file
func MakeRequest(ctx context.Context, method string, path string, body interface{}) (*http.Response, error) {
	client := http.Client{}
	if method == http.MethodGet {
		req, reqErr := http.NewRequestWithContext(ctx, method, path, nil)
		if reqErr != nil {
			return nil, reqErr
		}

		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")

		return client.Do(req)
	}

	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(encoded)
	req, reqErr := http.NewRequestWithContext(ctx, method, path, payload)
	if reqErr != nil {
		return nil, reqErr
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}
