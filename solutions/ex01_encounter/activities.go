package ex01_encounter

import (
	"context"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterWildPokemonActivity picks a random Pokemon from the wild.
func EncounterWildPokemonActivity(ctx context.Context) (pokemon.Pokemon, error) {
	p := pokemon.AllPokemon[rand.Intn(len(pokemon.AllPokemon))]
	return p, nil
}
