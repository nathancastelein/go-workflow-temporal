package helloworld

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func Helloworld(ctx workflow.Context, trainerName string, pokemonName string) (string, error) {
	retrypolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    0,
	}

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy:         retrypolicy,
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
