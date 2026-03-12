package ex02_capture

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

// ChoosePokemonActivity returns the trainer's signature Pokemon.
// TODO: Look up the trainer's Pokemon in pokemon.TrainerTeams.
// Return an error if the trainer is not found.
func ChoosePokemonActivity(ctx context.Context, trainerName string) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{}, nil
}

// WeakenActivity reduces the target's HP by attacker.HP/3, clamping to a minimum of 1.
// TODO: Calculate damage as attacker.HP / 3, subtract from target.HP,
// and clamp the result to a minimum of 1.
func WeakenActivity(ctx context.Context, attacker pokemon.Pokemon, target pokemon.Pokemon) (pokemon.Pokemon, error) {
	return pokemon.Pokemon{}, nil
}

// ThrowPokeballActivity attempts to capture the target Pokemon.
// TODO: Calculate capture probability as 1.0 - (target.HP / target.MaxHP).
// Use rand.Float64() to decide success. Return a CaptureResult.
func ThrowPokeballActivity(ctx context.Context, target pokemon.Pokemon) (pokemon.CaptureResult, error) {
	return pokemon.CaptureResult{}, nil
}
