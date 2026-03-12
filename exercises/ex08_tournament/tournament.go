package ex08_tournament

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// TournamentWorkflow orchestrates a full 8-trainer tournament with quarter-finals, semi-finals, and a final.
func TournamentWorkflow(ctx workflow.Context, trainers [8]string) (pokemon.TournamentResult, error) {
	// TODO:
	// 1. Set child workflow options:
	//    - ParentClosePolicy: enumspb.PARENT_CLOSE_POLICY_ABANDON
	//    - WorkflowExecutionTimeout: 1 minute
	//
	// 2. Quarter-finals — Launch 4 BattleWorkflow child workflows in parallel:
	//    - Use workflow.WithChildOptions(ctx, childOpts) to create a child context
	//    - Use workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, trainer1, trainer2) to get a future
	//    - Collect all 4 futures in a slice
	//
	// 3. Wait for quarter-final results:
	//    - For each future, call future.Get(ctx, &result)
	//    - If a child fails, the opponent advances by forfeit
	//    - Track winners in a [4]string array
	//
	// 4. Semi-finals — Launch 2 BattleWorkflow child workflows in parallel:
	//    - qfWinners[0] vs qfWinners[1], qfWinners[2] vs qfWinners[3]
	//
	// 5. Final — Launch 1 BattleWorkflow child workflow:
	//    - sfWinners[0] vs sfWinners[1]
	//
	// 6. Return TournamentResult with Bracket and Champion.
	return pokemon.TournamentResult{}, nil
}
