package ex01_encounter

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestEncounterWildPokemonActivity_ReturnsValidPokemon(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EncounterWildPokemonActivity)

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(EncounterWildPokemonActivity)

	// Assert
	require.NoError(t, err)

	var result pokemon.Pokemon
	require.NoError(t, encodedResult.Get(&result))
	assert.NotEmpty(t, result.Name)
	assert.Equal(t, result.HP, result.MaxHP)

	assert.Contains(t, pokemon.AllPokemon, result)
}
