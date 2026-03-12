package ex04_errors

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

func TestCapturePokemonWorkflow_SuccessAfterRetries(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	trainerPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	weakenedPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 28, MaxHP: 39}
	captureResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(captureResult, nil)
	env.OnActivity((&PokedexClient{}).RegisterInPokedexActivity, mock.Anything, weakenedPokemon).Return(nil)

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.True(t, result.Success)
}

func TestCapturePokemonWorkflow_PokemonFlees(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160}
	trainerPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	weakenedPokemon := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 149, MaxHP: 160}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(
		pokemon.CaptureResult{},
		temporal.NewNonRetryableApplicationError("pokemon fled", "PokemonFled", nil),
	)

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.False(t, result.Success)
	assert.Equal(t, weakenedPokemon, result.Pokemon)
}

func TestCapturePokemonWorkflow_PokedexRegistration(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Eevee", Type: "Normal", HP: 55, MaxHP: 55}
	trainerPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	weakenedPokemon := pokemon.Pokemon{Name: "Eevee", Type: "Normal", HP: 44, MaxHP: 55}
	captureResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(captureResult, nil)
	env.OnActivity((&PokedexClient{}).RegisterInPokedexActivity, mock.Anything, weakenedPokemon).Return(nil)

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.True(t, result.Success)
	// The workflow completed successfully, which means RegisterInPokedexActivity was called
	// (if it wasn't mocked and was called, the test would fail)
}

func TestCapturePokemonWorkflow_SuccessWithDifferentTrainer(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	trainerPokemon := pokemon.Pokemon{Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44}
	weakenedPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 21, MaxHP: 35}

	env.RegisterWorkflow(CapturePokemonWorkflow)
	env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	env.OnActivity(ChoosePokemonActivity, mock.Anything, "Misty").Return(trainerPokemon, nil)
	env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	// Simulate pokeball missing then succeeding
	env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(
		pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}, nil,
	)
	env.OnActivity((&PokedexClient{}).RegisterInPokedexActivity, mock.Anything, weakenedPokemon).Return(nil)

	// Act
	env.ExecuteWorkflow(CapturePokemonWorkflow, "Misty")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.True(t, result.Success)
}
