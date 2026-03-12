package ex05_testing

import (
	"testing"

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
	// TODO: Mock EncounterWildPokemonActivity to return Pikachu
	//   s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}, nil)
	// TODO: Mock DodgeCheckActivity to return false (no dodge)
	//   s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil)
	// TODO: Mock ChoosePokemonActivity to return Charmander
	//   s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(pokemon.Pokemon{Name: "Charmander", ...}, nil)
	// TODO: Mock WeakenActivity to return a weakened Pokemon
	//   s.env.OnActivity(WeakenActivity, mock.Anything, trainerPokemon, wildPokemon).Return(weakenedPokemon, nil)
	// TODO: Mock ThrowPokeballActivity to return CaptureResult{Success: true, ...}
	//   s.env.OnActivity(ThrowPokeballActivity, mock.Anything, weakenedPokemon).Return(pokemon.CaptureResult{Success: true, Pokemon: weakenedPokemon}, nil)

	// Act
	// TODO: Execute the workflow
	//   s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	// TODO: Assert workflow completed successfully
	//   s.True(s.env.IsWorkflowCompleted())
	//   s.NoError(s.env.GetWorkflowError())
	// TODO: Get and assert the result
	//   var result pokemon.CaptureResult
	//   s.NoError(s.env.GetWorkflowResult(&result))
	//   s.True(result.Success)
	s.Fail("implement this test")
}

func (s *WorkflowTestSuite) TestWorkflow_PokemonDodges() {
	// Arrange
	// TODO: Mock EncounterWildPokemonActivity to return a wild Pokemon
	//   s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil)
	// TODO: Mock DodgeCheckActivity to return true (Pokemon dodges!)
	//   s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(true, nil)
	// NOTE: Do NOT mock ChoosePokemonActivity, WeakenActivity, or ThrowPokeballActivity
	//   They should not be called when the Pokemon dodges!

	// Act
	// TODO: Execute the workflow
	//   s.env.ExecuteWorkflow(CapturePokemonWorkflow, "Ash")

	// Assert
	// TODO: Assert workflow completed and result shows failure
	//   s.True(s.env.IsWorkflowCompleted())
	//   s.NoError(s.env.GetWorkflowError())
	//   var result pokemon.CaptureResult
	//   s.NoError(s.env.GetWorkflowResult(&result))
	//   s.False(result.Success)
	s.Fail("implement this test")
}

func (s *WorkflowTestSuite) TestWorkflow_CaptureFailure() {
	// Arrange
	// TODO: Mock all activities for a full capture flow

	// Act
	// TODO: Execute the workflow

	// Assert
	// TODO: Make ThrowPokeballActivity return CaptureResult{Success: false, ...}
	// TODO: Assert result.Success == false
	s.Fail("implement this test")
}

func (s *WorkflowTestSuite) TestWorkflow_ActivityCallCount() {
	// Arrange
	// TODO: Mock all activities like in SuccessfulCapture
	// TODO: Add .Times(1) to each mock to verify each activity is called exactly once
	//   s.env.OnActivity(EncounterWildPokemonActivity, mock.Anything).Return(wildPokemon, nil).Times(1)
	//   s.env.OnActivity(DodgeCheckActivity, mock.Anything, wildPokemon).Return(false, nil).Times(1)
	//   ... and so on for all activities

	// Act
	// TODO: Execute the workflow

	// Assert
	// TODO: Assert workflow completed (AfterTest will verify mock expectations via AssertExpectations)
	s.Fail("implement this test")
}
