package ex09_tournament

import (
	"context"
	"fmt"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
)

// ChoosePokemonActivity returns the trainer's signature Pokemon from TrainerTeams.
func ChoosePokemonActivity(ctx context.Context, trainerName string) (pokemon.Pokemon, error) {
	p, ok := pokemon.TrainerTeams[trainerName]
	if !ok {
		return pokemon.Pokemon{}, fmt.Errorf("unknown trainer: %s", trainerName)
	}
	return p, nil
}
