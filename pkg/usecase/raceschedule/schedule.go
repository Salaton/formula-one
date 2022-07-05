package raceschedule

//go:generate mockgen -destination=mocks/mock.go -package=mock_raceschedule -source=schedule.go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/Salaton/formula-one/config"
	"github.com/Salaton/formula-one/pkg/domain"
)

type RaceSchedule interface {
	GetSeasonRaceSchedules(ctx context.Context, year int) (*domain.DataResponse, error)
}

type ScheduleDetails struct {
	config config.Configuration
}

func NewRaceScheduleImplementation() *ScheduleDetails {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}
	return &ScheduleDetails{
		config: *config,
	}
}

func (s ScheduleDetails) GetSeasonRaceSchedules(ctx context.Context, year int) (*domain.DataResponse, error) {
	var data *domain.DataResponse

	apiEndpoint := viper.GetString("ERGAST_API_ENDPOINT")
	apiResponseType := viper.GetString("ERGAST_API_RESPONSE_TYPE")

	urlPath := fmt.Sprintf(apiEndpoint, year, apiResponseType)
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
	encoded, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	payload := bytes.NewBuffer(encoded)
	req, err := http.NewRequestWithContext(ctx, method, path, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("an error occured while sending a HTTP request: %w", err)
	}

	return response, nil
}
