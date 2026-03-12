package ex09_queries

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterPokemonActivity picks a random Pokemon from the available list.
func EncounterPokemonActivity(ctx context.Context) (pokemon.Pokemon, error) {
	// TODO: Pick a random Pokemon from pokemon.AllPokemon
	// Hint: use rand.Intn(len(pokemon.AllPokemon)) to get a random index
	return pokemon.Pokemon{}, nil
}

// AttemptCaptureActivity attempts to capture a Pokemon.
// Capture succeeds if the Pokemon's HP is less than 100.
func AttemptCaptureActivity(ctx context.Context, p pokemon.Pokemon) (pokemon.CaptureResult, error) {
	// TODO: Check if p.HP < 100 for capture success
	// Return a CaptureResult with Success and the Pokemon
	return pokemon.CaptureResult{}, nil
}
