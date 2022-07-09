package graph

import (
	"context"

	"github.com/Salaton/formula-one/usecase"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	formulaone usecase.FormulaOne
}

func NewResolver(ctx context.Context, formulaone usecase.FormulaOne) *Resolver {
	return &Resolver{
		formulaone: formulaone,
	}
}
