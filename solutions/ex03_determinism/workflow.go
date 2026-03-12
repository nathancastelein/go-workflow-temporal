package ex03_determinism

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon.
// It introduces a dodge check: if the wild Pokemon dodges, the capture fails immediately.
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

	// Check if the Pokemon dodges — this is an activity to preserve determinism
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

	// Throw a Pokeball
	var result pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, ThrowPokeballActivity, weakened).Get(ctx, &result)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	return result, nil
}
