package usecase

import "github.com/Salaton/formula-one/pkg/usecase/raceschedule"

type FormulaOne struct {
	RaceSchedule raceschedule.RaceSchedule
}

func NewFormulaOneUseCase(
	RaceSchedule raceschedule.RaceSchedule,
) *FormulaOne {
	return &FormulaOne{
		RaceSchedule: RaceSchedule,
	}
}
