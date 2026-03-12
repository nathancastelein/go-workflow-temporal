package ex01_encounter

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestWildEncounterWorkflow_ReturnsValidPokemon(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	expectedPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	env.RegisterWorkflow(WildEncounterWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(expectedPokemon, nil)

	// Act
	env.ExecuteWorkflow(WildEncounterWorkflow)

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.Pokemon
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, expectedPokemon, result)
}
