package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Salaton/formula-one/pkg/domain"
	"github.com/Salaton/formula-one/pkg/presentation/graph/generated"
)

func (r *queryResolver) RaceSchedule(ctx context.Context, year int) (*domain.DataResponse, error) {
	return r.formulaone.RaceSchedule.GetSeasonRaceSchedules(ctx, year)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
