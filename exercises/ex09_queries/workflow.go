package ex09_queries

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// JourneyWorkflow simulates a trainer's Pokemon journey.
// It encounters and attempts to capture 3 Pokemon, updating the journey progress
// at each step. A query handler exposes the current progress.
func JourneyWorkflow(ctx workflow.Context, trainerName string) (pokemon.JourneyProgress, error) {
	// TODO: Create activity options with StartToCloseTimeout of 10 seconds
	//   ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second}
	//   ctx = workflow.WithActivityOptions(ctx, ao)

	// TODO: Initialize JourneyProgress with trainerName, empty CapturedPokemon slice,
	// and CurrentStatus "exploring"

	// TODO: Register a query handler for "progress" that returns the current JourneyProgress
	//   err := workflow.SetQueryHandler(ctx, "progress", func() (pokemon.JourneyProgress, error) {
	//       return progress, nil
	//   })

	// TODO: Loop 3 times, each iteration:
	//   1. Set CurrentStatus to "encountering"
	//   2. Call EncounterPokemonActivity to encounter a Pokemon
	//   3. Increment Encounters count
	//   4. Set CurrentStatus to "capturing"
	//   5. Call AttemptCaptureActivity with the encountered Pokemon
	//   6. If capture succeeded, append Pokemon to CapturedPokemon
	//   7. Set CurrentStatus back to "exploring"

	// TODO: Set CurrentStatus to "completed" and return the final progress

	_ = time.Second // remove when implementing
	return pokemon.JourneyProgress{}, nil
}
