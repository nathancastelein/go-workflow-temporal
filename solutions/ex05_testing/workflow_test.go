package ex05_testing

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestWorkflow_SuccessfulEvolution(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	raichu := pokemon.Pokemon{Name: "Raichu", Type: "Electric", HP: 60, MaxHP: 60}
	expectedResult := pokemon.EvolutionResult{Pokemon: raichu, Evolved: true, Trigger: "level-up"}

	env.OnActivity(FetchPokemonActivity, mock.Anything, "Pikachu").Return(pikachu, nil)
	env.OnActivity(CheckEvolutionActivity, mock.Anything, pikachu).Return("Raichu", nil)
	env.OnActivity(EvolvePokemonActivity, mock.Anything, pikachu, "Raichu").Return(expectedResult, nil)

	// Act
	env.ExecuteWorkflow(EvolvePokemonWorkflow, "Pikachu")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.EvolutionResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.True(t, result.Evolved)
	assert.Equal(t, "Raichu", result.Pokemon.Name)
	assert.Equal(t, "level-up", result.Trigger)
}
