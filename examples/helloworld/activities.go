package helloworld

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
)

func SayHelloToTrainer(ctx context.Context, trainerName string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SayHelloToTrainer", "name", trainerName)
	return fmt.Sprintf("Hello %s!", trainerName), nil
}

func SayHelloToPokemon(ctx context.Context, pokemonName string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SayHelloToPokemon", "name", pokemonName)
	return fmt.Sprintf("Hello %s!", pokemonName), nil
}
