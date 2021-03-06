package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Salaton/formula-one/graph/generated"
	model "github.com/Salaton/formula-one/models"
)

func (r *queryResolver) RaceSchedule(ctx context.Context, year int) (*model.DataResponse, error) {
	return r.formulaone.RaceSchedule.GetSeasonRaceSchedules(ctx, year)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
