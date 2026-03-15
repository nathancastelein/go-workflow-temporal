package ex04_errors

import (
	"testing"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func TestThrowPokeballActivity(t *testing.T) {
	t.Run("returns retryable error on first attempt", func(t *testing.T) {
		testSuite := &testsuite.WorkflowTestSuite{}
		activityEnv := testSuite.NewTestActivityEnvironment()
		activityEnv.RegisterActivity(ThrowPokeballActivity)

		target := pokemon.Pokemon{Name: "Charmander", Type: "Fire", HP: 1, MaxHP: 39}

		_, err := activityEnv.ExecuteActivity(ThrowPokeballActivity, target)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "pokeball missed")
	})

	// Note: non-retryable error (PokemonFled) and successful capture
	// depend on Attempt >= 2, which TestActivityEnvironment does not
	// simulate. These cases are covered by the workflow-level tests.
}

func TestRegisterInPokedexActivity(t *testing.T) {
	t.Run("returns retryable error on first attempt", func(t *testing.T) {
		testSuite := &testsuite.WorkflowTestSuite{}
		activityEnv := testSuite.NewTestActivityEnvironment()

		client := &PokedexClient{}
		activityEnv.RegisterActivity(client)

		p := pokemon.Pokemon{Name: "Pikachu", Type: "Electric", HP: 35, MaxHP: 35}

		_, err := activityEnv.ExecuteActivity(client.RegisterInPokedexActivity, p)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "pokedex API unavailable")
	})

	// Note: success after enough attempts depends on Attempt >= 3,
	// which TestActivityEnvironment does not simulate.
	// This case is covered by the workflow-level tests.
}
