package ex07_queries

import (
	"testing"
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestJourneyWorkflow_SuccessfulCapture(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(pikachu, nil)
	env.OnActivity(AttemptCaptureActivity, mock.Anything, pikachu).Return(
		pokemon.CaptureResult{Success: true, Pokemon: pikachu}, nil)

	// Act
	env.ExecuteWorkflow(JourneyWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.True(t, result.Success)
	assert.Equal(t, "Pikachu", result.Pokemon.Name)
}

func TestJourneyWorkflow_QueryReturnsProgress(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(pikachu, nil)
	env.OnActivity(AttemptCaptureActivity, mock.Anything, pikachu).Return(
		pokemon.CaptureResult{Success: true, Pokemon: pikachu}, nil)

	// Query during execution
	env.RegisterDelayedCallback(func() {
		encodedResult, err := env.QueryWorkflow("progress")
		require.NoError(t, err)
		var progress JourneyProgress
		require.NoError(t, encodedResult.Get(&progress))
		assert.Equal(t, "Ash", progress.TrainerName)
	}, time.Millisecond*1)

	// Act
	env.ExecuteWorkflow(JourneyWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	// Query after completion
	encodedResult, err := env.QueryWorkflow("progress")
	require.NoError(t, err)
	var progress JourneyProgress
	require.NoError(t, encodedResult.Get(&progress))
	assert.Equal(t, "captured", progress.CurrentStatus)
	assert.Equal(t, "Pikachu", progress.EncounteredPokemon.Name)
}
