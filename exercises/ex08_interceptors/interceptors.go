package ex08_interceptors

import (
	"context"

	"go.temporal.io/sdk/interceptor"
)

// TeamRocketReport records what Team Rocket observed about an activity execution.
type TeamRocketReport struct {
	ActivityName string
	Success      bool
}

// TeamRocketInterceptor spies on all activity executions and records reports.
// It implements both WorkerInterceptor (to wire itself up) and ActivityInboundInterceptor (to spy).
type TeamRocketInterceptor struct {
	interceptor.WorkerInterceptorBase
	interceptor.ActivityInboundInterceptorBase
	Reports []TeamRocketReport
}

func NewTeamRocketInterceptor() *TeamRocketInterceptor {
	return &TeamRocketInterceptor{}
}

// InterceptActivity inserts Team Rocket into the activity chain.
// Temporal calls this once at startup to build the interceptor chain.
func (tr *TeamRocketInterceptor) InterceptActivity(ctx context.Context, next interceptor.ActivityInboundInterceptor) interceptor.ActivityInboundInterceptor {
	tr.Next = next
	return tr
}

func (tr *TeamRocketInterceptor) ExecuteActivity(ctx context.Context, in *interceptor.ExecuteActivityInput) (interface{}, error) {
	// TODO: Delegate to tr.Next.ExecuteActivity(ctx, in)
	// TODO: Get the activity name using activity.GetInfo(ctx).ActivityType.Name
	// TODO: Append a TeamRocketReport with the activity name and whether it succeeded (err == nil)
	// TODO: Return the result and error from the delegation
	return tr.Next.ExecuteActivity(ctx, in)
}
