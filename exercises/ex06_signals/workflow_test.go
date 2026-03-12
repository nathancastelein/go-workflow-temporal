package ex06_signals

import (
	"testing"
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type EvolutionTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *EvolutionTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *EvolutionTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestEvolutionSuite(t *testing.T) {
	suite.Run(t, new(EvolutionTestSuite))
}

func (s *EvolutionTestSuite) TestEvolutionWorkflow_TimerTriggersEvolution() {
	// Arrange
	charmander := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	charmeleon := pokemon.Pokemon{Name: "Charmeleon", Type: "Fire", HP: 58, MaxHP: 58}

	s.env.OnActivity(EvolveActivity, mock.Anything, charmander).Return(charmeleon, nil)

	// Act
	s.env.ExecuteWorkflow(EvolutionWorkflow, charmander, 5*time.Second)

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.EvolutionResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.True(result.Evolved)
	s.Equal("timer", result.Trigger)
	s.Equal("Charmeleon", result.Pokemon.Name)
}

func (s *EvolutionTestSuite) TestEvolutionWorkflow_FeedSignalTriggersEvolution() {
	// Arrange
	charmander := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	charmeleon := pokemon.Pokemon{Name: "Charmeleon", Type: "Fire", HP: 58, MaxHP: 58}

	s.env.OnActivity(EvolveActivity, mock.Anything, charmander).Return(charmeleon, nil)

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow("feed", nil)
	}, 0)

	// Act
	s.env.ExecuteWorkflow(EvolutionWorkflow, charmander, 5*time.Second)

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.EvolutionResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.True(result.Evolved)
	s.Equal("feed", result.Trigger)
	s.Equal("Charmeleon", result.Pokemon.Name)
}

func (s *EvolutionTestSuite) TestEvolutionWorkflow_CancelSignalStopsEvolution() {
	// Arrange
	charmander := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}

	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow("cancel", nil)
	}, 0)

	// Act
	s.env.ExecuteWorkflow(EvolutionWorkflow, charmander, 5*time.Second)

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.EvolutionResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.False(result.Evolved)
	s.Equal("cancelled", result.Trigger)
	s.Equal("Charmander", result.Pokemon.Name)
}
