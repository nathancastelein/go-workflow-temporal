package ex05_testing

import (
	"testing"

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
	// TODO: Create a test activity environment
	//   activityEnv := s.NewTestActivityEnvironment()
	//   activityEnv.RegisterActivity(EncounterWildPokemonActivity)
	// TODO: Execute the activity
	//   encodedResult, err := activityEnv.ExecuteActivity(EncounterWildPokemonActivity)
	// TODO: Assert no error
	//   s.NoError(err)
	// TODO: Decode and assert the returned Pokemon has a non-empty name
	//   var result pokemon.Pokemon
	//   s.NoError(encodedResult.Get(&result))
	//   s.NotEmpty(result.Name)
	s.Fail("implement this test")
}
