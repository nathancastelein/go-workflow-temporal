package ex01_encounter

import (
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// WildEncounterWorkflow encounters a wild Pokemon.
// TODO: Create activity options with a StartToCloseTimeout of 10 seconds,
// then execute EncounterWildPokemonActivity and return the result.
func WildEncounterWorkflow(ctx workflow.Context) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{}, nil
}
