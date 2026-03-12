package ex07_interceptors

import (
	"context"
	"time"

	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/workflow"
)

// SpyReport records information about an activity execution.
type SpyReport struct {
	ActivityName string
	Duration     time.Duration
	Error        error
}

// SpyActivityInterceptor records activity execution details for inspection.
type SpyActivityInterceptor struct {
	interceptor.ActivityInboundInterceptorBase
	Reports []SpyReport
}

func (s *SpyActivityInterceptor) ExecuteActivity(ctx context.Context, in *interceptor.ExecuteActivityInput) (interface{}, error) {
	// TODO: Record the start time, delegate to s.Next.ExecuteActivity(ctx, in),
	// then append a SpyReport with the activity name, duration, and error.
	// Hint: Use activity.GetInfo(ctx) to get the activity name and time.Now()/time.Since() for timing.
	return s.Next.ExecuteActivity(ctx, in)
}

// TrainerCheckInterceptor validates that a trainer name argument is provided to the workflow.
type TrainerCheckInterceptor struct {
	interceptor.WorkflowInboundInterceptorBase
}

func (t *TrainerCheckInterceptor) ExecuteWorkflow(ctx workflow.Context, in *interceptor.ExecuteWorkflowInput) (interface{}, error) {
	// TODO: Check in.Args for a trainer name argument.
	// If in.Args is empty or the first argument is an empty string,
	// return fmt.Errorf("trainer name argument required").
	// Otherwise, delegate to t.Next.ExecuteWorkflow(ctx, in).
	return t.Next.ExecuteWorkflow(ctx, in)
}

// WorkerInterceptorFactory creates both the spy activity interceptor and the trainer check interceptor.
type WorkerInterceptorFactory struct {
	interceptor.WorkerInterceptorBase
	SpyInterceptor *SpyActivityInterceptor
}

func NewWorkerInterceptorFactory() *WorkerInterceptorFactory {
	return &WorkerInterceptorFactory{
		SpyInterceptor: &SpyActivityInterceptor{},
	}
}

func (w *WorkerInterceptorFactory) InterceptActivity(ctx context.Context, next interceptor.ActivityInboundInterceptor) interceptor.ActivityInboundInterceptor {
	// TODO: Set w.SpyInterceptor.Next = next and return w.SpyInterceptor
	return next
}

func (w *WorkerInterceptorFactory) InterceptWorkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor {
	// TODO: Create a TrainerCheckInterceptor, set its Next = next, and return it
	return next
}
