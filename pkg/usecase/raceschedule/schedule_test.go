package raceschedule_test

import (
	"context"
	"testing"

	"github.com/Salaton/formula-one/pkg/domain"
	"github.com/Salaton/formula-one/pkg/usecase/raceschedule"
	mock_raceschedule "github.com/Salaton/formula-one/pkg/usecase/raceschedule/mocks"
	"github.com/golang/mock/gomock"
)

func TestScheduleDetails_GetSeasonRaceSchedules(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRaceScheduleClient := mock_raceschedule.NewMockRaceSchedule(ctrl)

	// configure our mock to return mock data
	mockRaceScheduleClient.EXPECT().GetSeasonRaceSchedules(gomock.Any(), gomock.Any()).Return(&domain.DataResponse{
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

	// type args struct {
	// 	ctx  context.Context
	// 	year int
	// }
	// tests := []struct {
	// 	name    string
	// 	s       ScheduleDetails
	// 	args    args
	// 	want    *domain.DataResponse
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		s := ScheduleDetails{}
	// 		got, err := s.GetSeasonRaceSchedules(tt.args.ctx, tt.args.year)
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("ScheduleDetails.GetSeasonRaceSchedules() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("ScheduleDetails.GetSeasonRaceSchedules() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
