package ex10_search_attributes

import (
	"context"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterInRegionActivity returns a random Pokemon found in the given region.
func EncounterInRegionActivity(ctx context.Context, region string) (pokemon.Pokemon, error) {
	// TODO: Look up the region in pokemon.RegionPokemon
	// If the region doesn't exist, return an error
	// Otherwise, pick a random Pokemon from the list
	// Hint: use rand.Intn(len(pokemonList)) to get a random index
	return pokemon.Pokemon{}, nil
}

// AttemptCaptureActivity attempts to capture a Pokemon.
// Capture succeeds if the Pokemon's HP is less than 100.
func AttemptCaptureActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	// TODO: Return true if p.HP < 100, false otherwise
	return false, nil
}
