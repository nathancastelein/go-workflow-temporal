package ex03_determinism

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon.
// It introduces a flee check: if the wild Pokemon flees, the capture fails immediately.
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

	// Check if the Pokemon flees — this is an activity to preserve determinism
	var fled bool
	err = workflow.ExecuteActivity(ctx, FleeCheckActivity, wild).Get(ctx, &fled)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	if fled {
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
