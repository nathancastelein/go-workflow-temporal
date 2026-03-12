package ex02_capture

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the full capture sequence.
func CapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Encounter a wild Pokemon
	var wildPokemon pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, EncounterWildPokemonActivity).Get(ctx, &wildPokemon)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 2: Choose the trainer's Pokemon
	var trainerPokemon pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, ChoosePokemonActivity, trainerName).Get(ctx, &trainerPokemon)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 3: Weaken the wild Pokemon
	var weakenedPokemon pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, WeakenActivity, trainerPokemon, wildPokemon).Get(ctx, &weakenedPokemon)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 4: Throw a Pokeball
	var captureResult pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, ThrowPokeballActivity, weakenedPokemon).Get(ctx, &captureResult)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	return captureResult, nil
}
