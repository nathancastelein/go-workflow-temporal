package ex04_errors

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon with error handling and retry policies.
//
// The basic flow from Exercise 3 is already implemented below.
//
// TODO: Add error handling and retry policies:
//  1. For ThrowPokeballActivity, use separate activity options with a RetryPolicy:
//     - MaximumAttempts: 3
//     - InitialInterval: time.Second
//     - BackoffCoefficient: 2.0
//  2. If ThrowPokeballActivity fails (non-retryable error = pokemon fled),
//     return CaptureResult{Success: false, Pokemon: weakened}
//  3. If capture succeeds, call RegisterInPokedexActivity with its own options:
//     - StartToCloseTimeout: 5 * time.Second
//     - ScheduleToCloseTimeout: 30 * time.Second
//     - RetryPolicy: MaximumAttempts 5, InitialInterval 1s, BackoffCoefficient 2.0
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

	// Step 2: Check if the Pokemon flees
	var fled bool
	err = workflow.ExecuteActivity(ctx, FleeCheckActivity, wildPokemon).Get(ctx, &fled)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}
	if fled {
		return pokemon.CaptureResult{Success: false, Pokemon: wildPokemon}, nil
	}

	// Step 3: Choose the trainer's Pokemon
	var trainerPokemon pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, ChoosePokemonActivity, trainerName).Get(ctx, &trainerPokemon)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 4: Weaken the wild Pokemon
	var weakenedPokemon pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, WeakenActivity, trainerPokemon, wildPokemon).Get(ctx, &weakenedPokemon)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 5: Throw a Pokeball
	// TODO: Add a RetryPolicy to this activity call and handle the non-retryable error case
	var captureResult pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, ThrowPokeballActivity, weakenedPokemon).Get(ctx, &captureResult)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Step 6: Register in Pokedex
	// TODO: If capture succeeded, call RegisterInPokedexActivity with its own activity options

	return captureResult, nil
}
