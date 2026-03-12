package ex08_tournament

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// BattleWorkflow orchestrates a battle between two trainers.
func BattleWorkflow(ctx workflow.Context, trainer1, trainer2 string) (pokemon.BattleResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Each trainer chooses their Pokemon
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

	// Pokemon1 attacks Pokemon2
	var weakened2 pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, WeakenActivity, pokemon1, pokemon2).Get(ctx, &weakened2)
	if err != nil {
		return pokemon.BattleResult{}, err
	}

	// Pokemon2 attacks Pokemon1
	var weakened1 pokemon.Pokemon
	err = workflow.ExecuteActivity(ctx, WeakenActivity, pokemon2, pokemon1).Get(ctx, &weakened1)
	if err != nil {
		return pokemon.BattleResult{}, err
	}

	// Compare remaining HP to determine winner
	if weakened1.HP >= weakened2.HP {
		return pokemon.BattleResult{Winner: trainer1, Loser: trainer2}, nil
	}
	return pokemon.BattleResult{Winner: trainer2, Loser: trainer1}, nil
}
