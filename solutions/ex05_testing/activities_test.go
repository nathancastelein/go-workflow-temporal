package ex05_testing

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type ActivityTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestActivitySuite(t *testing.T) {
	suite.Run(t, new(ActivityTestSuite))
}

func (s *ActivityTestSuite) TestEncounterWildPokemonActivity_ReturnsValidPokemon() {
	// Arrange
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EncounterWildPokemonActivity)

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(EncounterWildPokemonActivity)

	// Assert
	s.NoError(err)
	var result pokemon.Pokemon
	s.NoError(encodedResult.Get(&result))
	s.NotEmpty(result.Name)
	s.NotEmpty(result.Type)
	s.Greater(result.HP, 0)
	s.Greater(result.MaxHP, 0)
}
