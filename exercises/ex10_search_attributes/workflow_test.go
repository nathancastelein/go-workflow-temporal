package ex10_search_attributes

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type ExpeditionTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *ExpeditionTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *ExpeditionTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestExpeditionSuite(t *testing.T) {
	suite.Run(t, new(ExpeditionTestSuite))
}

func (s *ExpeditionTestSuite) TestCaptureExpeditionWorkflow_SuccessfulCapture() {
	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	s.env.OnActivity(EncounterInRegionActivity, mock.Anything, "Kanto").Return(pikachu, nil)
	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, pikachu).Return(true, nil)

	s.env.ExecuteWorkflow(CaptureExpeditionWorkflow, "Ash", "Kanto")

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.ExpeditionResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("Ash", result.TrainerName)
	s.Equal("Kanto", result.Region)
	s.Equal("Pikachu", result.Pokemon.Name)
	s.True(result.Success)
}

func (s *ExpeditionTestSuite) TestCaptureExpeditionWorkflow_FailedCapture() {
	snorlax := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160}

	s.env.OnActivity(EncounterInRegionActivity, mock.Anything, "Hoenn").Return(snorlax, nil)
	s.env.OnActivity(AttemptCaptureActivity, mock.Anything, snorlax).Return(false, nil)

	s.env.ExecuteWorkflow(CaptureExpeditionWorkflow, "Misty", "Hoenn")

	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.ExpeditionResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("Misty", result.TrainerName)
	s.Equal("Hoenn", result.Region)
	s.Equal("Snorlax", result.Pokemon.Name)
	s.False(result.Success)
}
