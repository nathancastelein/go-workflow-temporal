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

	var pokemon1 pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, ChoosePokemonActivity, trainer1).Get(ctx, &pokemon1)
	if err != nil {
		return pokemon.BattleResult{}, err
	}

	var pokemon2 pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, ChoosePokemonActivity, trainer2).Get(ctx, &pokemon2)
	if err != nil {
		return pokemon.BattleResult{}, err
	}

	if pokemon1.HP >= pokemon2.HP {
		return pokemon.BattleResult{Winner: trainer1, Loser: trainer2}, nil
	}
	return pokemon.BattleResult{Winner: trainer2, Loser: trainer1}, nil
}
