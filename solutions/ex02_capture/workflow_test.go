package ex02_capture

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestCapturePokemonWorkflow_ReturnsResult(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Bulbasaur", Type: "Grass", HP: 45, MaxHP: 45}
	trainerPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	weakenedPokemon := pokemon.Pokemon{Name: "Bulbasaur", Type: "Grass", HP: 33, MaxHP: 45}
	expectedResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(expectedResult, nil)

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.Equal(t, expectedResult, result)
}
