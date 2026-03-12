package ex07_interceptors

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
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
	info := activity.GetInfo(ctx)
	start := time.Now()
	result, err := s.Next.ExecuteActivity(ctx, in)
	s.Reports = append(s.Reports, SpyReport{
		ActivityName: info.ActivityType.Name,
		Duration:     time.Since(start),
		Error:        err,
	})
	return result, err
}

// TrainerCheckInterceptor validates that a trainer name argument is provided to the workflow.
type TrainerCheckInterceptor struct {
	interceptor.WorkflowInboundInterceptorBase
}

func (t *TrainerCheckInterceptor) ExecuteWorkflow(ctx workflow.Context, in *interceptor.ExecuteWorkflowInput) (interface{}, error) {
	if len(in.Args) == 0 {
		return nil, fmt.Errorf("trainer name argument required")
	}
	trainerName, ok := in.Args[0].(string)
	if !ok || trainerName == "" {
		return nil, fmt.Errorf("trainer name argument required")
	}
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
	w.SpyInterceptor.Next = next
	return w.SpyInterceptor
}

func (w *WorkerInterceptorFactory) InterceptWorkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor {
	i := &TrainerCheckInterceptor{}
	i.Next = next
	return i
}
