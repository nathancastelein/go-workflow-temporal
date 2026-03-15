package ex05_testing

import (
	"context"
	"fmt"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// FetchPokemonActivity looks up a Pokemon by name from AllPokemon.
func FetchPokemonActivity(ctx context.Context, name string) (pokemon.Pokemon, error) {
	for _, p := range pokemon.AllPokemon {
		if p.Name == name {
			return p, nil
		}
	}
	return pokemon.Pokemon{}, fmt.Errorf("unknown pokemon: %s", name)
}

// CheckEvolutionActivity checks if the Pokemon has an evolution available.
// Returns the evolved form's name, or an error if the Pokemon cannot evolve.
func CheckEvolutionActivity(ctx context.Context, p pokemon.Pokemon) (string, error) {
	evolved, ok := pokemon.EvolutionMap[p.Name]
	if !ok {
		return "", fmt.Errorf("%s cannot evolve", p.Name)
	}
	return evolved.Name, nil
}

// EvolvePokemonActivity performs the evolution and returns the result.
func EvolvePokemonActivity(ctx context.Context, p pokemon.Pokemon, evolvedName string) (pokemon.EvolutionResult, error) {
	evolved, ok := pokemon.EvolutionMap[p.Name]
	if !ok {
		return pokemon.EvolutionResult{}, fmt.Errorf("no evolution found for %s", p.Name)
	}
	return pokemon.EvolutionResult{
		Pokemon: evolved,
		Evolved: true,
		Trigger: "level-up",
	}, nil
}
