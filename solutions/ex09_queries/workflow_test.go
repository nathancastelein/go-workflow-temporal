package ex09_queries

import (
	"testing"
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type JourneyTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *JourneyTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *JourneyTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestJourneySuite(t *testing.T) {
	suite.Run(t, new(JourneyTestSuite))
}

func (s *JourneyTestSuite) TestJourneyWorkflow_CompletesWithCaptures() {
	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	snorlax := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160}
	charmander := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}

	s.env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(pikachu, nil).Once()
	s.env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(snorlax, nil).Once()
	s.env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(charmander, nil).Once()

	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, pikachu).Return(
		pokemon.CaptureResult{Success: true, Pokemon: pikachu}, nil)
	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, snorlax).Return(
		pokemon.CaptureResult{Success: false, Pokemon: snorlax}, nil)
	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, charmander).Return(
		pokemon.CaptureResult{Success: true, Pokemon: charmander}, nil)

	s.env.ExecuteWorkflow(JourneyWorkflow, "Ash")

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.JourneyProgress
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("Ash", result.TrainerName)
	s.Equal(3, result.Encounters)
	s.Equal(2, len(result.CapturedPokemon))
	s.Equal("completed", result.CurrentStatus)
}

func (s *JourneyTestSuite) TestJourneyWorkflow_QueryReturnsProgress() {
	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	s.env.OnActivity(EncounterPokemonActivity, mock.Anything).Return(pikachu, nil)
	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, pikachu).Return(
		pokemon.CaptureResult{Success: true, Pokemon: pikachu}, nil)

	// Query after workflow completes to verify final state
	s.env.RegisterDelayedCallback(func() {
		encodedResult, err := s.env.QueryWorkflow("progress")
		s.NoError(err)
		var progress pokemon.JourneyProgress
		s.NoError(encodedResult.Get(&progress))
		s.Equal("Ash", progress.TrainerName)
	}, time.Millisecond*1)

	s.env.ExecuteWorkflow(JourneyWorkflow, "Ash")

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	// Query after completion
	encodedResult, err := s.env.QueryWorkflow("progress")
	s.NoError(err)
	var progress pokemon.JourneyProgress
	s.NoError(encodedResult.Get(&progress))
	s.Equal("completed", progress.CurrentStatus)
	s.Equal(3, progress.Encounters)
}
