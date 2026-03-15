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

// FleeCheckActivity returns true approximately 30% of the time.
// This randomness is safe because it runs inside an activity, not a workflow.
func FleeCheckActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	return rand.Float64() < 0.3, nil
}

// ThrowPokeballActivity attempts to capture the target Pokemon.
// TODO: Simulate unreliability based on the current attempt number:
//   - Attempt < 2: return a retryable error ("pokeball missed")
//   - Attempt >= 3 and capture would fail: return a non-retryable error (pokemon fled)
//   - Otherwise: return CaptureResult based on HP ratio probability (1.0 - target.HP/target.MaxHP)
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}

// PokedexClient simulates a client for an unreliable Pokedex registration API.
type PokedexClient struct{}

// RegisterInPokedexActivity registers a captured Pokemon in the Pokedex.
// TODO: Simulate an unreliable API:
//   - Attempt < 3: return a retryable error
//   - Attempt >= 3: succeed
func (c *PokedexClient) RegisterInPokedexActivity(ctx context.Context, p pokemon.Pokemon) error {
	return nil
}
