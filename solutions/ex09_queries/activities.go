package ex09_queries

import (
	"context"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterPokemonActivity picks a random Pokemon from the available list.
func EncounterPokemonActivity(ctx context.Context) (pokemon.Pokemon, error) {
	idx := rand.Intn(len(pokemon.AllPokemon))
	return pokemon.AllPokemon[idx], nil
}

// AttemptCaptureActivity attempts to capture a Pokemon.
// Capture succeeds if the Pokemon's HP is less than 100.
func AttemptCaptureActivity(ctx context.Context, p pokemon.Pokemon) (pokemon.CaptureResult, error) {
	success := p.HP < 100
	return pokemon.CaptureResult{Success: success, Pokemon: p}, nil
}
