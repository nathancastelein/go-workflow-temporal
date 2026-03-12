package ex04_errors

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon with error handling and retry policies.
func CapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Encounter a wild Pokemon
	var wild pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, EncounterWildPokemonActivity).Get(ctx, &wild)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Check if the Pokemon dodges
	var dodged bool
	err = workflow.ExecuteActivity(ctx, DodgeCheckActivity, wild).Get(ctx, &dodged)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	if dodged {
		return pokemon.CaptureResult{Success: false, Pokemon: wild}, nil
	}

	// Choose the trainer's Pokemon
	var chosen pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, ChoosePokemonActivity, trainerName).Get(ctx, &chosen)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Weaken the wild Pokemon
	var weakened pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, WeakenActivity, chosen, wild).Get(ctx, &weakened)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Throw a Pokeball with a retry policy
	throwCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    3,
			InitialInterval:    time.Second,
			BackoffCoefficient: 2.0,
		},
	})

	var result pokemon.CaptureResult
	err = workflow.ExecuteActivity(throwCtx, ThrowPokeballActivity, weakened).Get(ctx, &result)
	if err != nil {
		// If the pokeball throw fails (non-retryable = pokemon fled), return failure
		return pokemon.CaptureResult{Success: false, Pokemon: weakened}, nil
	}

	// If capture succeeded, register in the Pokedex
	if result.Success {
		pokedexCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout:    5 * time.Second,
			ScheduleToCloseTimeout: 30 * time.Second,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts:    5,
				InitialInterval:    time.Second,
				BackoffCoefficient: 2.0,
			},
		})

		pokedexClient := &PokedexClient{}
		err = workflow.ExecuteActivity(pokedexCtx, pokedexClient.RegisterInPokedexActivity, result.Pokemon).Get(ctx, nil)
		if err != nil {
			return pokemon.CaptureResult{}, err
		}
	}

	return result, nil
}
