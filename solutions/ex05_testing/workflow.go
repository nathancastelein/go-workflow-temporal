package ex05_testing

import (
	"time"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/workflow"
)

// EvolvePokemonWorkflow orchestrates the evolution of a Pokemon.
// It fetches the Pokemon, checks if it can evolve, and performs the evolution.
func EvolvePokemonWorkflow(ctx workflow.Context, pokemonName string) (pokemon.EvolutionResult, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Step 1: Fetch the Pokemon by name
	var p pokemon.Pokemon
	err := workflow.ExecuteActivity(ctx, FetchPokemonActivity, pokemonName).Get(ctx, &p)
	if err != nil {
		return pokemon.EvolutionResult{}, err
	}

	// Step 2: Check if the Pokemon can evolve
	var evolvedName string
	err = workflow.ExecuteActivity(ctx, CheckEvolutionActivity, p).Get(ctx, &evolvedName)
	if err != nil {
		return pokemon.EvolutionResult{}, err
	}

	// Step 3: Perform the evolution
	var result pokemon.EvolutionResult
	err = workflow.ExecuteActivity(ctx, EvolvePokemonActivity, p, evolvedName).Get(ctx, &result)
	if err != nil {
		return pokemon.EvolutionResult{}, err
	}

	return result, nil
}
