package helloworld

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func Helloworld(ctx workflow.Context, trainerName string, pokemonName string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result, finalResult string
	err := workflow.ExecuteActivity(ctx, SayHelloToTrainer, trainerName).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	finalResult += result

	err = workflow.ExecuteActivity(ctx, SayHelloToPokemon, pokemonName).Get(ctx, &result)
	if err != nil {
		return "", err
	}
	finalResult += " " + result

	return finalResult, nil
}
