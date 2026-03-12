package ex04_errors

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
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
// This randomness is safe because it runs inside an activity, not a workflow.
func DodgeCheckActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	return rand.Float64() < 0.3, nil
}

// ThrowPokeballActivity attempts to capture the target Pokemon.
// It simulates unreliability using activity.GetInfo(ctx).Attempt:
//   - If attempt < 2: return a retryable error ("pokeball missed")
//   - If attempt >= 3 and capture would fail: return a non-retryable error
//     using temporal.NewNonRetryableApplicationError("pokemon fled", "PokemonFled", nil)
//   - Otherwise: return CaptureResult based on HP ratio probability
//
// TODO: Implement the attempt-based retry logic described above.
// Hint: Use activity.GetInfo(ctx).Attempt to get the current attempt number.
// Hint: Use temporal.NewNonRetryableApplicationError for non-retryable errors.
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}

// PokedexClient simulates a client for an unreliable Pokedex registration API.
type PokedexClient struct{}

// RegisterInPokedexActivity registers a captured Pokemon in the Pokedex.
// It simulates an unreliable API using activity.GetInfo(ctx).Attempt:
//   - If attempt < 3: return a retryable error
//   - If attempt >= 3: succeed
//
// TODO: Implement the attempt-based retry logic.
// This activity uses a struct receiver so it must be registered with:
//
//	w.RegisterActivity(&PokedexClient{})
func (c *PokedexClient) RegisterInPokedexActivity(ctx context.Context, p pokemon.Pokemon) error {
	return nil
}
