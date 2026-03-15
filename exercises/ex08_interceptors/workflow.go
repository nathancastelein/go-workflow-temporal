package ex08_interceptors

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CatchPokemonWorkflow encounters a wild Pokemon and throws a Pokeball.
func CatchPokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var wild pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, EncounterActivity).Get(ctx, &wild)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	var result pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, ThrowPokeballActivity, wild).Get(ctx, &result)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	return result, nil
}
