package raceschedule_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	model "github.com/Salaton/formula-one/models"
	"github.com/Salaton/formula-one/usecase/raceschedule"
	mock_raceschedule "github.com/Salaton/formula-one/usecase/raceschedule/mocks"
)

func TestScheduleDetails_GetSeasonRaceSchedules(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRaceScheduleClient := mock_raceschedule.NewMockRaceSchedule(ctrl)

	// configure our mock to return mock data
	mockRaceScheduleClient.EXPECT().GetSeasonRaceSchedules(context.Background(), 1990).Return(&model.DataResponse{
		MRData: model.MRData{
			RaceTable: &model.RaceTable{
				Season: "1990",
				Races: []*model.Races{{
					Season: "1990",
					Round:  "1",
				}},
			},
		},
	}, nil).AnyTimes()

	s := raceschedule.ScheduleDetails{}
	s.GetSeasonRaceSchedules(context.Background(), 1990)
}
