package ex04_errors

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// CapturePokemonWorkflow orchestrates the capture of a wild Pokemon with error handling and retry policies.
//
// TODO:
//  1. Create default activity options with StartToCloseTimeout of 10 seconds
//  2. Call EncounterWildPokemonActivity, DodgeCheckActivity, ChoosePokemonActivity, WeakenActivity
//  3. For ThrowPokeballActivity, use separate activity options with a RetryPolicy:
//     - MaximumAttempts: 3
//     - InitialInterval: time.Second
//     - BackoffCoefficient: 2.0
//  4. If ThrowPokeballActivity fails (non-retryable error = pokemon fled),
//     return CaptureResult{Success: false, Pokemon: weakened}
//  5. If capture succeeds, call RegisterInPokedexActivity with its own options:
//     - StartToCloseTimeout: 5 * time.Second
//     - ScheduleToCloseTimeout: 30 * time.Second
//     - RetryPolicy: MaximumAttempts 5, InitialInterval 1s, BackoffCoefficient 2.0
//  6. Return the CaptureResult
func CapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}
