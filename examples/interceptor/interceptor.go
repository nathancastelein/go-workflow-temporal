package interceptor

import (
	"context"
	"log/slog"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/interceptor"
)

// LoggingInterceptor logs the activity name and whether it succeeded.
type LoggingInterceptor struct {
	interceptor.ActivityInboundInterceptorBase
}

func (l *LoggingInterceptor) ExecuteActivity(ctx context.Context, in *interceptor.ExecuteActivityInput) (interface{}, error) {
	name := activity.GetInfo(ctx).ActivityType.Name
	slog.Info("activity started", "activity", name)

	result, err := l.Next.ExecuteActivity(ctx, in)

	slog.Info("activity finished", "activity", name, "success", err == nil)
	return result, err
}

// LoggingWorkerInterceptor is the factory that wires the LoggingInterceptor.
type LoggingWorkerInterceptor struct {
	interceptor.WorkerInterceptorBase
}

func (f *LoggingWorkerInterceptor) InterceptActivity(
	ctx context.Context,
	next interceptor.ActivityInboundInterceptor,
) interceptor.ActivityInboundInterceptor {
	i := &LoggingInterceptor{}
	i.Next = next
	return i
}
