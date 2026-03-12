package ex05_testing

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *WorkflowTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *WorkflowTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestWorkflowSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (s *WorkflowTestSuite) TestWorkflow_SuccessfulCapture() {
	// Arrange
	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	trainerPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	weakenedPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 22, MaxHP: 35}
	expectedResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	s.env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	s.env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(expectedResult, nil)

	// Act
	s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.CaptureResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.True(result.Success)
}

func (s *WorkflowTestSuite) TestWorkflow_PokemonDodges() {
	// Arrange
	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(true, nil)

	// Act
	s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.CaptureResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.False(result.Success)
	s.Equal("Pikachu", result.Pokemon.Name)
}

func (s *WorkflowTestSuite) TestWorkflow_CaptureFailure() {
	// Arrange
	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	trainerPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	weakenedPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 22, MaxHP: 35}
	failedResult := pokemon.CaptureResult{Success: false, Pokemon: weakenedPokemon}

	s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil)
	s.env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	s.env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(failedResult, nil)

	// Act
	s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.CaptureResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.False(result.Success)
}

func (s *WorkflowTestSuite) TestWorkflow_ActivityCallCount() {
	// Arrange
	wildPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	trainerPokemon := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 39, MaxHP: 39}
	weakenedPokemon := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 22, MaxHP: 35}
	expectedResult := pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}

	s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil).Times(1)
	s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil).Times(1)
	s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(trainerPokemon, nil).Times(1)
	s.env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil).Times(1)
	s.env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(expectedResult, nil).Times(1)

	// Act
	s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
}
