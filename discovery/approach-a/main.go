package main

import (
	"context"
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

type Event struct {
	Type    string
	Pokemon Pokemon
	Trainer Trainer
}

func main() {
	// Create communication channels between workers
	statusCh := make(chan Event, 1)
	combatCh := make(chan Event, 1)
	pokeballCh := make(chan Event, 1)
	doneCh := make(chan Event, 1)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Worker 1: listens for encounters, paralyzes the pokemon, forwards to combat
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-statusCh:
				slog.Info("paralyzing pokemon", "pokemon", event.Pokemon.Name)
				fmt.Printf("[Worker 1] Paralyzing %s...\n", event.Pokemon.Name)
				combatCh <- Event{Type: "paralyzed", Pokemon: event.Pokemon, Trainer: event.Trainer}
			}
		}
	}()

	// Worker 2: listens for paralyzed pokemon, attacks it, forwards to pokeball
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-combatCh:
				slog.Info("attacking pokemon", "pokemon", event.Pokemon.Name)
				fmt.Printf("[Worker 2] Attacking %s...\n", event.Pokemon.Name)
				pokeballCh <- Event{Type: "weakened", Pokemon: event.Pokemon, Trainer: event.Trainer}
			}
		}
	}()

	// Worker 3: listens for weakened pokemon, throws pokeball, signals done
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-pokeballCh:
				slog.Info("throwing pokeball", "pokemon", event.Pokemon.Name, "trainer", event.Trainer.Name)
				fmt.Printf("[Worker 3] %s throws a pokeball at %s...\n", event.Trainer.Name, event.Pokemon.Name)
				fmt.Printf("[Worker 3] %s captured!\n", event.Pokemon.Name)
				doneCh <- Event{Type: "captured", Pokemon: event.Pokemon, Trainer: event.Trainer}
			}
		}
	}()

	// Start the workflow by sending the initial event
	fmt.Println("=== Starting capture workflow ===")
	statusCh <- Event{
		Type:    "encountered",
		Pokemon: Pokemon{Name: "Mewtwo"},
		Trainer: Trainer{Name: "Sacha"},
	}

	// Wait for completion
	result := <-doneCh
	fmt.Printf("\n=== Workflow complete: %s captured %s ===\n", result.Trainer.Name, result.Pokemon.Name)
}
