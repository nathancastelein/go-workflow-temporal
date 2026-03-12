package ex01_encounter

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterWildPokemonActivity picks a random Pokemon from the wild.
// TODO: Pick a random Pokemon from pokemon.AllPokemon and return it.
func EncounterWildPokemonActivity(ctx context.Context) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{}, nil
}
