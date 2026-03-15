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
	_ = childOpts

	// TODO: Semi-finals — launch 2 parallel BattleWorkflow child workflows:
	//   - lorelei vs bruno
	//   - agatha vs lance
	// Use workflow.WithChildOptions and workflow.ExecuteChildWorkflow.
	// Collect the 2 winners using future.Get().

	// TODO: Final — launch 1 BattleWorkflow child workflow with the 2 winners.

	// TODO: Return TournamentResult with Champion.

	return pokemon.TournamentResult{}, nil
}
