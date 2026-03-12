package main

import (
	"fmt"
	"log/slog"
)

// Domain types (self-contained)
type Pokemon struct {
	Name string
}

type Trainer struct {
	Name string
}

type CaptureResult struct {
	Pokemon Pokemon
	Trainer Trainer
	Success bool
}

// Steps — each step is a function that takes the current state and returns an updated state
func paralyze(pokemon Pokemon) Pokemon {
	slog.Info("paralyzing pokemon", "pokemon", pokemon.Name)
	fmt.Printf("[Step 1] Paralyzing %s...\n", pokemon.Name)
	return pokemon
}

func attack(pokemon Pokemon) Pokemon {
	slog.Info("attacking pokemon", "pokemon", pokemon.Name)
	fmt.Printf("[Step 2] Attacking %s...\n", pokemon.Name)
	return pokemon
}

func throwPokeball(trainer Trainer, pokemon Pokemon) CaptureResult {
	slog.Info("throwing pokeball", "pokemon", pokemon.Name, "trainer", trainer.Name)
	fmt.Printf("[Step 3] %s throws a pokeball at %s...\n", trainer.Name, pokemon.Name)
	fmt.Printf("[Step 3] %s captured!\n", pokemon.Name)
	return CaptureResult{Pokemon: pokemon, Trainer: trainer, Success: true}
}

// capturePokemon drives the entire workflow from a single function
func capturePokemon(trainer Trainer, pokemon Pokemon) CaptureResult {
	fmt.Println("--- Running capture workflow ---")

	// Step 1: Paralyze
	pokemon = paralyze(pokemon)

	// Step 2: Attack
	pokemon = attack(pokemon)

	// Step 3: Throw pokeball
	result := throwPokeball(trainer, pokemon)

	fmt.Println("--- Workflow complete ---")
	return result
}

func main() {
	trainer := Trainer{Name: "Sacha"}
	pokemon := Pokemon{Name: "Mewtwo"}

	fmt.Println("=== Starting capture workflow ===")
	result := capturePokemon(trainer, pokemon)
	fmt.Printf("\n=== Workflow complete: %s captured %s ===\n", result.Trainer.Name, result.Pokemon.Name)
}
