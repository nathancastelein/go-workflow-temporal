package ex02_capture

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the full capture sequence.
// TODO:
// 1. Create activity options with a StartToCloseTimeout of 10 seconds
// 2. Call EncounterWildPokemonActivity to encounter a wild Pokemon
// 3. Call ChoosePokemonActivity with the trainerName to get the trainer's Pokemon
// 4. Call WeakenActivity with the trainer's Pokemon and the wild Pokemon
// 5. Call ThrowPokeballActivity with the weakened Pokemon
// 6. Return the CaptureResult
func CapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}
