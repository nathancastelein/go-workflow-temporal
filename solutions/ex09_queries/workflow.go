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
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	progress := pokemon.JourneyProgress{
		TrainerName:     trainerName,
		Encounters:      0,
		CapturedPokemon: []pokemon.Pokemon{},
		CurrentStatus:   "exploring",
	}

	err := workflow.SetQueryHandler(ctx, "progress", func() (pokemon.JourneyProgress, error) {
		return progress, nil
	})
	if err != nil {
		return progress, err
	}

	for i := 0; i < 3; i++ {
		// Encounter
		progress.CurrentStatus = "encountering"
		var encountered pokemon.Pokemon
		err := workflow.ExecuteActivity(ctx, EncounterPokemonActivity).Get(ctx, &encountered)
		if err != nil {
			return progress, err
		}
		progress.Encounters++

		// Capture attempt
		progress.CurrentStatus = "capturing"
		var result pokemon.CaptureResult
		err = workflow.ExecuteActivity(ctx, AttemptCaptureActivity, encountered).Get(ctx, &result)
		if err != nil {
			return progress, err
		}
		if result.Success {
			progress.CapturedPokemon = append(progress.CapturedPokemon, result.Pokemon)
		}

		progress.CurrentStatus = "exploring"
	}

	progress.CurrentStatus = "completed"
	return progress, nil
}
