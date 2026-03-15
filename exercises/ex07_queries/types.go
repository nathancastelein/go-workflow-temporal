package ex07_queries

import "github.com/nathancastelein/go-workflow-temporal/pokemon"

// JourneyProgress tracks the current state of a trainer's Pokemon journey.
type JourneyProgress struct {
	TrainerName        string
	EncounteredPokemon pokemon.Pokemon
	CurrentStatus      string // "exploring", "encountering", "capturing", "captured", "escaped"
}
