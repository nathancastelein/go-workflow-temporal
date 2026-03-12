package ex02_capture

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestWeakenActivity_ReducesHP(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(WeakenActivity)

	attacker := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	target := pokemon.Pokemon{Name: "Bulbasaur", Type: "Grass", HP: 45, MaxHP: 45}

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(WeakenActivity, attacker, target)

	// Assert
	require.NoError(t, err)

	var result pokemon.Pokemon
	require.NoError(t, encodedResult.Get(&result))
	assert.Less(t, result.HP, 45)
	assert.GreaterOrEqual(t, result.HP, 1)
}

func TestWeakenActivity_ClampsToMinimum1(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(WeakenActivity)

	attacker := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 300, MaxHP: 300}
	target := pokemon.Pokemon{Name: "Gastly", Type: "Ghost", HP: 10, MaxHP: 30}

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(WeakenActivity, attacker, target)

	// Assert
	require.NoError(t, err)

	var result pokemon.Pokemon
	require.NoError(t, encodedResult.Get(&result))
	assert.Equal(t, 1, result.HP)
}

func TestChoosePokemonActivity_ReturnsTrainerPokemon(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(ChoosePokemonActivity)

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(ChoosePokemonActivity, "Ash")

	// Assert
	require.NoError(t, err)

	var result pokemon.Pokemon
	require.NoError(t, encodedResult.Get(&result))
	assert.Equal(t, "Pikachu", result.Name)
}

func TestChoosePokemonActivity_ErrorOnUnknownTrainer(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(ChoosePokemonActivity)

	// Act
	_, err := activityEnv.ExecuteActivity(ChoosePokemonActivity, "Unknown")

	// Assert
	assert.Error(t, err)
}
