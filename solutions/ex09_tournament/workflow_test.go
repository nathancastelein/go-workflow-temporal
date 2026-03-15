package ex09_tournament

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestBattleWorkflow(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	squirtle := pokemon.Pokemon{Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44}
	machop := pokemon.Pokemon{Name: "Machop", Type: "Fighting", HP: 70, MaxHP: 70}

	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Lorelei").Return(squirtle, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Bruno").Return(machop, nil)

	// Act
	env.ExecuteWorkflow(BattleWorkflow, "Lorelei", "Bruno")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.BattleResult
	require.NoError(t, env.GetWorkflowResult(&result))
	// Machop (70 HP) > Squirtle (44 HP), so Bruno wins
	assert.Equal(t, "Bruno", result.Winner)
	assert.Equal(t, "Lorelei", result.Loser)
}

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
