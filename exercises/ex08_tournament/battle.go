package ex08_tournament

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// BattleWorkflow orchestrates a battle between two trainers.
func BattleWorkflow(ctx workflow.Context, trainer1, trainer2 string) (pokemon.BattleResult, error) {
	// TODO:
	// 1. Set activity options with StartToCloseTimeout of 10 seconds.
	// 2. Call ChoosePokemonActivity for each trainer.
	// 3. Call WeakenActivity: pokemon1 attacks pokemon2, then pokemon2 attacks pokemon1.
	// 4. Compare remaining HP: if weakened1.HP >= weakened2.HP, trainer1 wins.
	// 5. Return BattleResult with Winner and Loser.
	return pokemon.BattleResult{}, nil
}
