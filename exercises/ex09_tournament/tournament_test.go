package ex09_tournament

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)



func TestTournamentWorkflow(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Semi-finals
	env.OnWorkflow(BattleWorkflow, mock.Anything, "Lorelei", "Bruno").Return(pokemon.BattleResult{Winner: "Bruno", Loser: "Lorelei"}, nil)
	env.OnWorkflow(BattleWorkflow, mock.Anything, "Agatha", "Lance").Return(pokemon.BattleResult{Winner: "Lance", Loser: "Agatha"}, nil)

	// Final
	env.OnWorkflow(BattleWorkflow, mock.Anything, "Bruno", "Lance").Return(pokemon.BattleResult{Winner: "Bruno", Loser: "Lance"}, nil)

	// Act
	env.ExecuteWorkflow(TournamentWorkflow, "Lorelei", "Bruno", "Agatha", "Lance")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.TournamentResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, "Bruno", result.Champion)
}
