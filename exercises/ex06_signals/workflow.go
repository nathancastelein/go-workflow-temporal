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
	// TODO: Create activity options with StartToCloseTimeout of 10 seconds
	//   ao := workflow.ActivityOptions{StartToCloseTimeout: 10 * time.Second}
	//   ctx = workflow.WithActivityOptions(ctx, ao)

	// TODO: Create signal channels for "feed" and "cancel"
	//   feedCh := workflow.GetSignalChannel(ctx, "feed")
	//   cancelCh := workflow.GetSignalChannel(ctx, "cancel")

	// TODO: Create a timer that fires after the given duration
	//   timerFuture := workflow.NewTimer(ctx, duration)

	// TODO: Create a selector to wait for the first event
	//   sel := workflow.NewSelector(ctx)

	// TODO: Add timer case - when timer fires, call EvolveActivity and return with Trigger "timer"
	//   sel.AddFuture(timerFuture, func(f workflow.Future) {
	//       var evolved pokemon.Pokemon
	//       workflow.ExecuteActivity(ctx, EvolveActivity, p).Get(ctx, &evolved)
	//       result = pokemon.EvolutionResult{Pokemon: evolved, Evolved: true, Trigger: "timer"}
	//   })

	// TODO: Add feed signal case - when feed signal received, call EvolveActivity and return with Trigger "feed"
	//   sel.AddReceive(feedCh, func(ch workflow.ReceiveChannel, more bool) {
	//       ch.Receive(ctx, nil)
	//       ...evolve and set result with Trigger "feed"
	//   })

	// TODO: Add cancel signal case - return Pokemon unchanged with Trigger "cancelled"
	//   sel.AddReceive(cancelCh, func(ch workflow.ReceiveChannel, more bool) {
	//       ch.Receive(ctx, nil)
	//       result = pokemon.EvolutionResult{Pokemon: p, Evolved: false, Trigger: "cancelled"}
	//   })

	// TODO: Call sel.Select(ctx) to wait for the first event

	return pokemon.EvolutionResult{}, nil
}
