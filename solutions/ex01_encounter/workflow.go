package ex01_encounter

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// WildEncounterWorkflow encounters a wild Pokemon.
func WildEncounterWorkflow(ctx workflow.Context) (pokemon.Pokemon, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, EncounterWildPokemonActivity).Get(ctx, &result)
	if err != nil {
		return pokemon.Pokemon{}, err
	}

	return result, nil
}
