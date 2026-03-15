package ex09_tournament

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
)

// TournamentWorkflow orchestrates a 4-trainer tournament with semi-finals and a final.
func TournamentWorkflow(ctx workflow.Context, lorelei, bruno, agatha, lance string) (pokemon.TournamentResult, error) {
	childOpts := workflow.ChildWorkflowOptions{
		ParentClosePolicy:        enumspb.PARENT_CLOSE_POLICY_ABANDON,
		WorkflowExecutionTimeout: 1 * time.Minute,
	}

	// Semi-finals: 2 parallel battles
	sf1 := workflow.ExecuteChildWorkflow(workflow.WithChildOptions(ctx, childOpts), BattleWorkflow, lorelei, bruno)
	sf2 := workflow.ExecuteChildWorkflow(workflow.WithChildOptions(ctx, childOpts), BattleWorkflow, agatha, lance)

	var result1, result2 pokemon.BattleResult
	if err := sf1.Get(ctx, &result1); err != nil {
		return pokemon.TournamentResult{}, err
	}
	if err := sf2.Get(ctx, &result2); err != nil {
		return pokemon.TournamentResult{}, err
	}

	// Final
	var finalResult pokemon.BattleResult
	err := workflow.ExecuteChildWorkflow(
		workflow.WithChildOptions(ctx, childOpts), BattleWorkflow, result1.Winner, result2.Winner,
	).Get(ctx, &finalResult)
	if err != nil {
		return pokemon.TournamentResult{}, err
	}

	return pokemon.TournamentResult{
		Champion: finalResult.Winner,
	}, nil
}
