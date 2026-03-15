package ex03_determinism

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestCapturePokemonWorkflow_PokemonFlees(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(FleeCheckActivity, mock.Anything, wildPokemon).Return(true, nil)
	// WeakenActivity and ThrowPokeballActivity are NOT mocked — if the workflow
	// calls them, the test will fail, proving the flee short-circuits correctly.

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.False(t, result.Success)
	assert.Equal(t, wildPokemon, result.Pokemon)
}

func TestCapturePokemonWorkflow_NoFlee_SuccessfulCapture(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	trainerPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	weakenedPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 28, MaxHP: 39}
	expectedResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(FleeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
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
	assert.True(t, result.Success)
	assert.Equal(t, expectedResult, result)
}
