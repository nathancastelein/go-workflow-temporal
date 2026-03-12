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
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	feedCh := workflow.GetSignalChannel(ctx, "feed")
	cancelCh := workflow.GetSignalChannel(ctx, "cancel")

	timerFuture := workflow.NewTimer(ctx, duration)

	var result pokemon.EvolutionResult

	sel := workflow.NewSelector(ctx)

	sel.AddFuture(timerFuture, func(f workflow.Future) {
		var evolved pokemon.Pokemon
		err := workflow.ExecuteActivity(ctx, EvolveActivity, p).Get(ctx, &evolved)
		if err != nil {
			result = pokemon.EvolutionResult{Pokemon: p, Evolved: false, Trigger: "timer"}
			return
		}
		result = pokemon.EvolutionResult{Pokemon: evolved, Evolved: true, Trigger: "timer"}
	})

	sel.AddReceive(feedCh, func(ch workflow.ReceiveChannel, more bool) {
		ch.Receive(ctx, nil)
		var evolved pokemon.Pokemon
		err := workflow.ExecuteActivity(ctx, EvolveActivity, p).Get(ctx, &evolved)
		if err != nil {
			result = pokemon.EvolutionResult{Pokemon: p, Evolved: false, Trigger: "feed"}
			return
		}
		result = pokemon.EvolutionResult{Pokemon: evolved, Evolved: true, Trigger: "feed"}
	})

	sel.AddReceive(cancelCh, func(ch workflow.ReceiveChannel, more bool) {
		ch.Receive(ctx, nil)
		result = pokemon.EvolutionResult{Pokemon: p, Evolved: false, Trigger: "cancelled"}
	})

	sel.Select(ctx)

	return result, nil
}
