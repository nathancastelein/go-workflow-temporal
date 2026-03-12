package ex10_search_attributes

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// EncounterInRegionActivity returns a random Pokemon found in the given region.
func EncounterInRegionActivity(ctx context.Context, region string) (pokemon.Pokemon, error) {
	pokemonList, ok := pokemon.RegionPokemon[region]
	if !ok {
		return pokemon.Pokemon{}, fmt.Errorf("unknown region: %s", region)
	}
	idx := rand.Intn(len(pokemonList))
	return pokemonList[idx], nil
}

// AttemptCaptureActivity attempts to capture a Pokemon.
// Capture succeeds if the Pokemon's HP is less than 100.
func AttemptCaptureActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	return p.HP < 100, nil
}
