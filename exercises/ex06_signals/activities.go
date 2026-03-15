package ex06_signals

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EvolveActivity looks up the evolved form of a Pokemon in pokemon.EvolutionMap.
// If the Pokemon has an evolution, return it. Otherwise, return the Pokemon unchanged.
// TODO: Implement the evolution lookup.
func EvolveActivity(ctx context.Context, p pokemon.Pokemon) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{}, nil
}
