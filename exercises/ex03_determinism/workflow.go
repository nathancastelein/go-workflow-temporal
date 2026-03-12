package ex03_determinism

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon.
// It introduces a dodge check: if the wild Pokemon dodges, the capture fails immediately.
//
// TODO:
// 1. Create activity options with StartToCloseTimeout of 10 seconds
// 2. Call EncounterWildPokemonActivity to encounter a wild Pokemon
// 3. Call DodgeCheckActivity to check if the Pokemon dodges
// 4. If dodged, return CaptureResult{Success: false, Pokemon: wild}
// 5. If not dodged, call ChoosePokemonActivity, WeakenActivity, ThrowPokeballActivity
// 6. Return the CaptureResult
func CapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}
