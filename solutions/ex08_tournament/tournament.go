package ex08_tournament

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
)

// TournamentWorkflow orchestrates a full 8-trainer tournament with quarter-finals, semi-finals, and a final.
func TournamentWorkflow(ctx workflow.Context, trainers [8]string) (pokemon.TournamentResult, error) {
	childOpts := workflow.ChildWorkflowOptions{
		ParentClosePolicy:        enumspb.PARENT_CLOSE_POLICY_ABANDON,
		WorkflowExecutionTimeout: 1 * time.Minute,
	}

	var bracket [][]string

	// Quarter-finals: 4 parallel battles
	var quarterFinals []workflow.ChildWorkflowFuture
	for i := 0; i < 8; i += 2 {
		childCtx := workflow.WithChildOptions(ctx, childOpts)
		future := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, trainers[i], trainers[i+1])
		quarterFinals = append(quarterFinals, future)
	}

	var qfWinners [4]string
	var qfRound []string
	for i, future := range quarterFinals {
		var result pokemon.BattleResult
		if err := future.Get(ctx, &result); err != nil {
			// Opponent advances by forfeit
			if i*2+1 < 8 {
				qfWinners[i] = trainers[i*2+1]
			}
		} else {
			qfWinners[i] = result.Winner
		}
		qfRound = append(qfRound, qfWinners[i])
	}
	bracket = append(bracket, qfRound)

	// Semi-finals: 2 parallel battles
	var semiFinals []workflow.ChildWorkflowFuture
	for i := 0; i < 4; i += 2 {
		childCtx := workflow.WithChildOptions(ctx, childOpts)
		future := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, qfWinners[i], qfWinners[i+1])
		semiFinals = append(semiFinals, future)
	}

	var sfWinners [2]string
	var sfRound []string
	for i, future := range semiFinals {
		var result pokemon.BattleResult
		if err := future.Get(ctx, &result); err != nil {
			sfWinners[i] = qfWinners[i*2+1]
		} else {
			sfWinners[i] = result.Winner
		}
		sfRound = append(sfRound, sfWinners[i])
	}
	bracket = append(bracket, sfRound)

	// Final
	childCtx := workflow.WithChildOptions(ctx, childOpts)
	finalFuture := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, sfWinners[0], sfWinners[1])

	var finalResult pokemon.BattleResult
	var champion string
	if err := finalFuture.Get(ctx, &finalResult); err != nil {
		champion = sfWinners[1]
	} else {
		champion = finalResult.Winner
	}
	bracket = append(bracket, []string{champion})

	return pokemon.TournamentResult{
		Bracket:  bracket,
		Champion: champion,
	}, nil
}
