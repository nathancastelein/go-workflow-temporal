package ex08_interceptors

import (
	"context"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterActivity returns a random wild Pokemon.
func EncounterActivity(ctx context.Context) (pokemon.Pokemon, error) {
	return pokemon.AllPokemon[rand.Intn(len(pokemon.AllPokemon))], nil
}

// ThrowPokeballActivity attempts to catch the target Pokemon (30% success rate).
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	success := rand.Float64() < 0.3
	return pokemon.CaptureResult{Success: success, Pokemon: target}, nil
}
