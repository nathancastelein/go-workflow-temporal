package ex07_interceptors

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// SimpleWorkflow takes a trainer name, calls SimpleActivity and returns its result.
func SimpleWorkflow(ctx workflow.Context, trainerName string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, SimpleActivity).Get(ctx, &result)
	if err != nil {
		return "", err
	}

	return result, nil
}
