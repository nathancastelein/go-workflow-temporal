package ex03_determinism

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestFleeCheckActivity_ReturnsBool(t *testing.T) {
	// Arrange
	testSuite := &testsuite.WorkflowTestSuite{}
	activityEnv := testSuite.NewTestActivityEnvironment()
	activityEnv.RegisterActivity(FleeCheckActivity)
	p := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

	// Act
	encodedResult, err := activityEnv.ExecuteActivity(FleeCheckActivity, p)

	// Assert
	require.NoError(t, err)
	var result bool
	require.NoError(t, encodedResult.Get(&result))
	assert.IsType(t, true, result)
}
