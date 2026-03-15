package ex06_signals

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// EvolutionWorkflow waits for either a timer, a feed signal, or a cancel signal.
// - Timer fires: evolve the Pokemon and return with Trigger "timer"
// - Feed signal: evolve the Pokemon and return with Trigger "feed"
// - Cancel signal: return the Pokemon unchanged with Trigger "cancelled"
func EvolutionWorkflow(ctx workflow.Context, p pokemon.Pokemon, duration time.Duration) (pokemon.EvolutionResult, error) {
	// TODO: Set up activity options
	// TODO: Create signal channels for "feed" and "cancel"
	// TODO: Create a timer with the given duration
	// TODO: Use a selector to wait for the first event:
	//   - Timer fires: call EvolveActivity, return with Trigger "timer"
	//   - Feed signal received: call EvolveActivity, return with Trigger "feed"
	//   - Cancel signal received: return Pokemon unchanged with Trigger "cancelled"

	return pokemon.EvolutionResult{}, nil
}
