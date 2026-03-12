package ex08_tournament

import (
	"fmt"
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type BattleTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *BattleTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *BattleTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestBattleSuite(t *testing.T) {
	suite.Run(t, new(BattleTestSuite))
}

func (s *BattleTestSuite) TestBattleWorkflow_ReturnsWinner() {
	// Arrange
	// Ash's Pikachu (HP 35) vs Misty's Squirtle (HP 44)
	pikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}
	squirtle := pokemon.Pokemon{Name: "Squirtle", Type: "Water", HP: 44, MaxHP: 44}

	// After Pikachu attacks Squirtle: 44 - 35/3 = 44 - 11 = 33
	weakenedSquirtle := pokemon.Pokemon{Name: "Squirtle", Type: "Water", HP: 33, MaxHP: 44}
	// After Squirtle attacks Pikachu: 35 - 44/3 = 35 - 14 = 21
	weakenedPikachu := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 21, MaxHP: 35}

	s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Ash").Return(pikachu, nil)
	s.env.OnActivity(ChoosePokemonActivity, mock.Anything, "Misty").Return(squirtle, nil)
	s.env.OnActivity(WeakenActivity, mock.Anything, pikachu, squirtle).Return(weakenedSquirtle, nil)
	s.env.OnActivity(WeakenActivity, mock.Anything, squirtle, pikachu).Return(weakenedPikachu, nil)

	// Act
	s.env.ExecuteWorkflow(BattleWorkflow, "Ash", "Misty")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.BattleResult
	s.NoError(s.env.GetWorkflowResult(&result))
	// weakenedPikachu.HP (21) < weakenedSquirtle.HP (33), so Misty wins
	s.Equal("Misty", result.Winner)
	s.Equal("Ash", result.Loser)
}

// Tournament tests

type TournamentTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env *testsuite.TestWorkflowEnvironment
}

func (s *TournamentTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
}

func (s *TournamentTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestTournamentSuite(t *testing.T) {
	suite.Run(t, new(TournamentTestSuite))
}

func (s *TournamentTestSuite) TestTournamentWorkflow_FullBracket() {
	// Arrange
	trainers := [8]string{"Ash", "Misty", "Brock", "Gary", "Jessie", "James", "Sabrina", "Giovanni"}

	// Quarter-finals
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Ash", "Misty").Return(pokemon.BattleResult{Winner: "Ash", Loser: "Misty"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Brock", "Gary").Return(pokemon.BattleResult{Winner: "Gary", Loser: "Brock"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Jessie", "James").Return(pokemon.BattleResult{Winner: "Jessie", Loser: "James"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Sabrina", "Giovanni").Return(pokemon.BattleResult{Winner: "Giovanni", Loser: "Sabrina"}, nil)

	// Semi-finals
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Ash", "Gary").Return(pokemon.BattleResult{Winner: "Ash", Loser: "Gary"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Jessie", "Giovanni").Return(pokemon.BattleResult{Winner: "Giovanni", Loser: "Jessie"}, nil)

	// Final
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Ash", "Giovanni").Return(pokemon.BattleResult{Winner: "Ash", Loser: "Giovanni"}, nil)

	// Act
	s.env.ExecuteWorkflow(TournamentWorkflow, trainers)

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.TournamentResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("Ash", result.Champion)
	s.Len(result.Bracket, 3)
	s.Len(result.Bracket[0], 4) // Quarter-final winners
	s.Len(result.Bracket[1], 2) // Semi-final winners
	s.Len(result.Bracket[2], 1) // Final winner
}

func (s *TournamentTestSuite) TestTournamentWorkflow_ChildFailure() {
	// Arrange
	trainers := [8]string{"Ash", "Misty", "Brock", "Gary", "Jessie", "James", "Sabrina", "Giovanni"}

	// First quarter-final fails — Misty advances by forfeit
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Ash", "Misty").Return(pokemon.BattleResult{}, fmt.Errorf("battle failed"))
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Brock", "Gary").Return(pokemon.BattleResult{Winner: "Gary", Loser: "Brock"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Jessie", "James").Return(pokemon.BattleResult{Winner: "Jessie", Loser: "James"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Sabrina", "Giovanni").Return(pokemon.BattleResult{Winner: "Giovanni", Loser: "Sabrina"}, nil)

	// Semi-finals: Misty (forfeit winner) vs Gary
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Misty", "Gary").Return(pokemon.BattleResult{Winner: "Gary", Loser: "Misty"}, nil)
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Jessie", "Giovanni").Return(pokemon.BattleResult{Winner: "Giovanni", Loser: "Jessie"}, nil)

	// Final
	s.env.OnWorkflow(BattleWorkflow, mock.Anything, "Gary", "Giovanni").Return(pokemon.BattleResult{Winner: "Giovanni", Loser: "Gary"}, nil)

	// Act
	s.env.ExecuteWorkflow(TournamentWorkflow, trainers)

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result pokemon.TournamentResult
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("Giovanni", result.Champion)
	// Misty should have advanced by forfeit in the first round
	s.Equal("Misty", result.Bracket[0][0])
}
