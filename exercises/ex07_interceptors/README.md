# Exercise 7 — Team Rocket is Watching

## Concepts

In this exercise, you will learn about **interceptors** in Temporal — a mechanism for cross-cutting concerns like logging, metrics, authentication, and input validation.

### Interceptor Interfaces

Temporal's Go SDK provides an interceptor chain pattern:

- **`interceptor.WorkerInterceptor`** — A factory interface that creates interceptors for activities and workflows. It has two methods:
  - `InterceptActivity(ctx context.Context, next interceptor.ActivityInboundInterceptor) interceptor.ActivityInboundInterceptor`
  - `InterceptWorkflow(ctx workflow.Context, next interceptor.WorkflowInboundInterceptor) interceptor.WorkflowInboundInterceptor`

- **`interceptor.ActivityInboundInterceptor`** — Wraps activity execution. Embed `interceptor.ActivityInboundInterceptorBase` and override `ExecuteActivity` to add behavior before/after the activity runs.

- **`interceptor.WorkflowInboundInterceptor`** — Wraps workflow execution. Embed `interceptor.WorkflowInboundInterceptorBase` and override `ExecuteWorkflow` to add behavior before/after the workflow runs.

### Wiring Interceptors

Register interceptors when creating the worker:

```go
factory := NewWorkerInterceptorFactory()
w := worker.New(c, taskQueue, worker.Options{
    Interceptors: []interceptor.WorkerInterceptor{factory},
})
```

### Interceptor Input

In a workflow interceptor's `ExecuteWorkflow`, you receive `*interceptor.ExecuteWorkflowInput` which contains:
- `Args []interface{}` — the decoded arguments passed to the workflow

You can inspect or validate these arguments before delegating to the next interceptor in the chain.

For activity interceptors, use `activity.GetInfo(ctx)` to access activity metadata like the activity type name.

## What to implement

### Spy Activity Interceptor (`interceptors.go`)

Implement `SpyActivityInterceptor.ExecuteActivity` — it should:
1. Record the start time
2. Delegate to `s.Next.ExecuteActivity(ctx, in)`
3. Append a `SpyReport` with the activity name (use `activity.GetInfo(ctx).ActivityType.Name`), duration, and any error

### Trainer Check Interceptor (`interceptors.go`)

Implement `TrainerCheckInterceptor.ExecuteWorkflow` — it should:
1. Check `in.Args` for a trainer name argument
2. If `in.Args` is empty or the first argument is an empty string, return `fmt.Errorf("trainer name argument required")`
3. Otherwise, delegate to `t.Next.ExecuteWorkflow(ctx, in)`

### Worker Interceptor Factory (`interceptors.go`)

Implement the factory methods:
- `InterceptActivity`: Wire up the `SpyActivityInterceptor` by setting its `Next` field
- `InterceptWorkflow`: Create and wire up a `TrainerCheckInterceptor`

## How to test

```bash
go test ./exercises/ex07_interceptors/...
```

## How to run

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex07_interceptors/worker/
   ```
