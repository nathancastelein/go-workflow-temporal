package ex08_interceptors

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

func TestCatchWorkflow_TeamRocketSpies(t *testing.T) {
	// Arrange
	teamRocket := NewTeamRocketInterceptor()
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.SetWorkerOptions(worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{teamRocket},
	})

	// Register real activities so interceptors are invoked
	env.RegisterActivity(EncounterActivity)
	env.RegisterActivity(ThrowPokeballActivity)

	// Act
	env.ExecuteWorkflow(CatchPokemonWorkflow, "Ash")

	// Assert
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var result pokemon.CaptureResult
	require.NoError(t, env.GetWorkflowResult(&result))
	assert.NotEmpty(t, result.Pokemon.Name)

	// Team Rocket should have observed both activities
	require.Len(t, teamRocket.Reports, 2)
	assert.Equal(t, "EncounterActivity", teamRocket.Reports[0].ActivityName)
	assert.True(t, teamRocket.Reports[0].Success)
	assert.Equal(t, "ThrowPokeballActivity", teamRocket.Reports[1].ActivityName)
	assert.True(t, teamRocket.Reports[1].Success)
}
