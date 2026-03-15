package ex07_queries

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// JourneyWorkflow simulates a trainer encountering and trying to capture a Pokemon.
// A query handler exposes the current progress.
func JourneyWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	progress := JourneyProgress{
		TrainerName:   trainerName,
		CurrentStatus: "exploring",
	}

	err := workflow.SetQueryHandler(ctx, "progress", func() (JourneyProgress, error) {
		return progress, nil
	})
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Encounter a wild Pokemon
	progress.CurrentStatus = "encountering"
	var encountered pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, EncounterPokemonActivity).Get(ctx, &encountered)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Attempt capture
	progress.CurrentStatus = "capturing"
	var result pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, AttemptCaptureActivity, encountered).Get(ctx, &result)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	if result.Success {
		progress.CurrentStatus = "captured"
		progress.EncounteredPokemon = result.Pokemon
	} else {
		progress.CurrentStatus = "escaped"
	}

	// simulate a long process
	_ = workflow.NewTimer(ctx, 5*time.Minute).Get(ctx, nil)

	return result, nil
}
