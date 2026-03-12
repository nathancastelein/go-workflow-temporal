package ex05_testing

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the full capture sequence with dodge check and retry.
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

	// Step 2: Check if the wild Pokemon dodges
	var dodged bool
	err = workflow.ExecuteActivity(ctx, DodgeCheckActivity, wildPokemon).Get(ctx, &dodged)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}
	if dodged {
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

	// Step 5: Throw a Pokeball with retry policy
	throwCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    3,
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
		},
	})

	var captureResult pokemon.CaptureResult
	err = workflow.ExecuteActivity(throwCtx, ThrowPokeballActivity, weakenedPokemon).Get(ctx, &captureResult)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	return captureResult, nil
}
