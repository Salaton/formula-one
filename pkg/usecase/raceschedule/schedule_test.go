package raceschedule_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/Salaton/formula-one/pkg/domain"
	"github.com/Salaton/formula-one/pkg/usecase/raceschedule"
	mock_raceschedule "github.com/Salaton/formula-one/pkg/usecase/raceschedule/mocks"
)

func TestScheduleDetails_GetSeasonRaceSchedules(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRaceScheduleClient := mock_raceschedule.NewMockRaceSchedule(ctrl)

	// configure our mock to return mock data
	mockRaceScheduleClient.EXPECT().GetSeasonRaceSchedules(context.Background(), 1990).Return(&domain.DataResponse{
		MRData: domain.MRData{
			RaceTable: &domain.RaceTable{
				Season: "1990",
				Races: []*domain.Races{{
					Season: "1990",
					Round:  "1",
				}},
			},
		},
	}, nil).AnyTimes()

	s := raceschedule.ScheduleDetails{}
	s.GetSeasonRaceSchedules(context.Background(), 1990)
}
