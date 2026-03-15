package ex08_interceptors

import (
	"context"

	"go.temporal.io/sdk/activity"
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

func (tr *TeamRocketInterceptor) InterceptActivity(ctx context.Context, next interceptor.ActivityInboundInterceptor) interceptor.ActivityInboundInterceptor {
	tr.Next = next
	return tr
}

func (tr *TeamRocketInterceptor) ExecuteActivity(ctx context.Context, in *interceptor.ExecuteActivityInput) (interface{}, error) {
	result, err := tr.Next.ExecuteActivity(ctx, in)

	name := activity.GetInfo(ctx).ActivityType.Name
	tr.Reports = append(tr.Reports, TeamRocketReport{
		ActivityName: name,
		Success:      err == nil,
	})

	return result, err
}
