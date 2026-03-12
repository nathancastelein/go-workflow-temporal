package ex06_signals

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

func (s *ActivityTestSuite) TestEvolveActivity_CharmandertToCharmeleon() {
	// Arrange
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EvolveActivity)
	charmander := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(EvolveActivity, charmander)

	// Assert
	s.NoError(err)
	var result pokemon.Pokemon
	s.NoError(encodedResult.Get(&result))
	s.Equal("Charmeleon", result.Name)
	s.Equal("Fire", result.Type)
	s.Equal(58, result.HP)
}

func (s *ActivityTestSuite) TestEvolveActivity_SnorlaxStaysSame() {
	// Arrange
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EvolveActivity)
	snorlax := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160}

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(EvolveActivity, snorlax)

	// Assert
	s.NoError(err)
	var result pokemon.Pokemon
	s.NoError(encodedResult.Get(&result))
	s.Equal("Snorlax", result.Name)
	s.Equal(160, result.HP)
}
