package ex09_tournament

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// BattleWorkflow orchestrates a battle between two trainers.
// Each trainer chooses their Pokemon, then the one with higher HP wins.
func BattleWorkflow(ctx workflow.Context, trainer1, trainer2 string) (pokemon.BattleResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// TODO: Call ChoosePokemonActivity for trainer1
	// TODO: Call ChoosePokemonActivity for trainer2
	// TODO: Compare HP — higher HP wins
	// TODO: Return a BattleResult with Winner and Loser

	return pokemon.BattleResult{}, nil
}
