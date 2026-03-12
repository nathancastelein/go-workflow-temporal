package ex08_tournament

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

// WeakenActivity reduces the target's HP by attacker.HP/3, clamping to a minimum of 1.
func WeakenActivity(ctx context.Context, attacker, target pokemon.Pokemon) (pokemon.Pokemon, error) {
	damage := attacker.HP / 3
	target.HP -= damage
	if target.HP < 1 {
		target.HP = 1
	}
	return target, nil
}
