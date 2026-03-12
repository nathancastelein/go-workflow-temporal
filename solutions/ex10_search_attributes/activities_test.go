package ex10_search_attributes

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

func (s *ActivityTestSuite) TestEncounterInRegionActivity_Kanto() {
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EncounterInRegionActivity)

	encodedResult, err := activityEnv.ExecuteActivity(EncounterInRegionActivity, "Kanto")

	s.NoError(err)
	var result pokemon.Pokemon
	s.NoError(encodedResult.Get(&result))
	s.NotEmpty(result.Name)
	s.Greater(result.MaxHP, 0)
}

func (s *ActivityTestSuite) TestEncounterInRegionActivity_UnknownRegion() {
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(EncounterInRegionActivity)

	_, err := activityEnv.ExecuteActivity(EncounterInRegionActivity, "Unknown")

	s.Error(err)
}

func (s *ActivityTestSuite) TestAttemptCaptureActivity_LowHPSucceeds() {
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(AttemptCaptureActivity)
	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	encodedResult, err := activityEnv.ExecuteActivity(AttemptCaptureActivity, pikachu)

	s.NoError(err)
	var result bool
	s.NoError(encodedResult.Get(&result))
	s.True(result)
}

func (s *ActivityTestSuite) TestAttemptCaptureActivity_HighHPFails() {
	activityEnv := s.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(AttemptCaptureActivity)
	snorlax := pokemon.Pokemon{Name: "Snorlax", Type: "Normal", HP: 160, MaxHP: 160}

	encodedResult, err := activityEnv.ExecuteActivity(AttemptCaptureActivity, snorlax)

	s.NoError(err)
	var result bool
	s.NoError(encodedResult.Get(&result))
	s.False(result)
}
