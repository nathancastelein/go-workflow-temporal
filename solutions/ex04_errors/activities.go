package ex04_errors

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
)

// EncounterWildPokemonActivity picks a random Pokemon from the wild.
func EncounterWildPokemonActivity(ctx context.Context) (pokemon.Pokemon, error) {
	p := pokemon.AllPokemon[rand.Intn(len(pokemon.AllPokemon))]
	return p, nil
}

// ChoosePokemonActivity returns the trainer's signature Pokemon.
func ChoosePokemonActivity(ctx context.Context, trainerName string) (pokemon.Pokemon, error) {
	p, ok := pokemon.TrainerTeams[trainerName]
	if !ok {
		return pokemon.Pokemon{}, fmt.Errorf("unknown trainer: %s", trainerName)
	}
	return p, nil
}

// WeakenActivity reduces the target's HP by attacker.HP/3, clamping to a minimum of 1.
func WeakenActivity(ctx context.Context, attacker pokemon.Pokemon, target pokemon.Pokemon) (pokemon.Pokemon, error) {
	damage := attacker.HP / 3
	target.HP -= damage
	if target.HP < 1 {
		target.HP = 1
	}
	return target, nil
}

// DodgeCheckActivity returns true approximately 30% of the time.
func DodgeCheckActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	return rand.Float64() < 0.3, nil
}

// ThrowPokeballActivity attempts to capture the target Pokemon.
// It simulates unreliability: the first attempt always fails with a retryable error.
// If the capture would fail on attempt >= 3, it returns a non-retryable error (pokemon fled).
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	info := activity.GetInfo(ctx)

	// Simulate the pokeball missing on early attempts
	if info.Attempt < 2 {
		return pokemon.CaptureResult{}, fmt.Errorf("pokeball missed")
	}

	probability := 1.0 - (float64(target.HP) / float64(target.MaxHP))
	success := rand.Float64() < probability

	if !success && info.Attempt >= 3 {
		return pokemon.CaptureResult{}, temporal.NewNonRetryableApplicationError("pokemon fled", "PokemonFled", nil)
	}

	return pokemon.CaptureResult{
		Success: success,
		Pokemon: target,
	}, nil
}

// PokedexClient simulates a client for an unreliable Pokedex registration API.
type PokedexClient struct{}

// RegisterInPokedexActivity registers a captured Pokemon in the Pokedex.
// It simulates an unreliable API that fails on early attempts.
func (c *PokedexClient) RegisterInPokedexActivity(ctx context.Context, p pokemon.Pokemon) error {
	info := activity.GetInfo(ctx)

	// Simulate API failures on early attempts
	if info.Attempt < 3 {
		return fmt.Errorf("pokedex API unavailable (attempt %d)", info.Attempt)
	}

	return nil
}
