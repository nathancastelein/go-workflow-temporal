package ex06_signals

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EvolveActivity looks up the evolved form of a Pokemon.
// If the Pokemon has an evolution, it returns the evolved form.
// If not, it returns the same Pokemon unchanged.
func EvolveActivity(ctx context.Context, p pokemon.Pokemon) (pokemon.Pokemon, error) {
	// TODO: Look up p.Name in pokemon.EvolutionMap
	// If found, return the evolved Pokemon
	// If not found, return the same Pokemon unchanged
	return pokemon.Pokemon{}, nil
}
