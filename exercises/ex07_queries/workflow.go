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

	// TODO: Create a JourneyProgress variable to track the journey state.
	// Initialize it with the trainer name and status "exploring".

	// TODO: Register a query handler named "progress" that returns the current JourneyProgress.
	// Use workflow.SetQueryHandler with a func() (JourneyProgress, error) callback.

	// Encounter a wild Pokemon
	// TODO: Update the progress status to "encountering"
	var encountered pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, EncounterPokemonActivity).Get(ctx, &encountered)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// Attempt capture
	// TODO: Update the progress status to "capturing"
	var result pokemon.CaptureResult
	err = workflow.ExecuteActivity(ctx, AttemptCaptureActivity, encountered).Get(ctx, &result)
	if err != nil {
		return pokemon.CaptureResult{}, err
	}

	// TODO: Update the progress:
	// - If capture succeeded: set status to "captured" and store the encountered Pokemon
	// - If capture failed: set status to "escaped"

	// simulate a long process
	_ = workflow.NewTimer(ctx, 5*time.Minute).Get(ctx, nil)

	return result, nil
}
