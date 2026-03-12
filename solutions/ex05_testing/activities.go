package ex05_testing

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

// DodgeCheckActivity determines if the wild Pokemon dodges (~30% chance).
func DodgeCheckActivity(ctx context.Context, p pokemon.Pokemon) (bool, error) {
	return rand.Float64() < 0.3, nil
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

// ThrowPokeballActivity attempts to capture the target Pokemon.
// Capture probability = 1.0 - (target.HP / target.MaxHP).
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	probability := 1.0 - (float64(target.HP) / float64(target.MaxHP))
	success := rand.Float64() < probability
	return pokemon.CaptureResult{
		Success: success,
		Pokemon: target,
	}, nil
}
