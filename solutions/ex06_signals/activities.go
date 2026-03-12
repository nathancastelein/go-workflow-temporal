package ex06_signals

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EvolveActivity looks up the evolved form of a Pokemon.
// If the Pokemon has an evolution, it returns the evolved form.
// If not, it returns the same Pokemon unchanged.
func EvolveActivity(ctx context.Context, p pokemon.Pokemon) (pokemon.Pokemon, error) {
	evolved, ok := pokemon.EvolutionMap[p.Name]
	if ok {
		return evolved, nil
	}
	return p, nil
}
