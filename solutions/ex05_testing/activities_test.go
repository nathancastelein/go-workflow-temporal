package ex05_testing

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestFetchPokemonActivity_KnownPokemon(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(FetchPokemonActivity)

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(FetchPokemonActivity, "Pikachu")

	// Assert
	require.NoError(t, err)
	var result pokemon.Pokemon
	require.NoError(t, encodedResult.Get(&result))
	assert.Equal(t, "Pikachu", result.Name)
	assert.Equal(t, "Electric", result.Type)
	assert.Equal(t, 35, result.HP)
	assert.Equal(t, 35, result.MaxHP)
}

func TestFetchPokemonActivity_UnknownPokemon(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(FetchPokemonActivity)

	// Act
	_, err := activityEnv.ExecuteActivity(FetchPokemonActivity, "MissingNo")

	// Assert
	assert.Error(t, err)
}
