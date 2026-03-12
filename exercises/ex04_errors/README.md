# Exercise 4 — The Pokeball Missed! (Error Handling)

## Concepts

In this exercise, you will learn about **error handling and retry policies** in Temporal.

### RetryPolicy

Temporal can automatically retry failed activities. You configure this with `temporal.RetryPolicy`:

```go
ao := workflow.ActivityOptions{
    StartToCloseTimeout: 10 * time.Second,
    RetryPolicy: &temporal.RetryPolicy{
        MaximumAttempts:    3,              // Max number of attempts (including first)
        InitialInterval:    time.Second,    // Wait before first retry
        BackoffCoefficient: 2.0,           // Multiply interval by this each retry
    },
}
```

- **`MaximumAttempts`**: Total number of attempts. 3 means: first try + 2 retries.
- **`InitialInterval`**: How long to wait before the first retry.
- **`BackoffCoefficient`**: Each subsequent retry waits longer: 1s, 2s, 4s, ...

### Non-Retryable Errors

Sometimes an error is permanent — retrying won't help. Use `temporal.NewNonRetryableApplicationError` to stop retries immediately:

```go
return temporal.NewNonRetryableApplicationError("pokemon fled", "PokemonFled", nil)
```

Temporal will not retry this error, even if `MaximumAttempts` hasn't been reached.

### StartToCloseTimeout vs ScheduleToCloseTimeout

- **`StartToCloseTimeout`**: Maximum time for a single attempt of an activity.
- **`ScheduleToCloseTimeout`**: Maximum total time from when the activity is scheduled until it completes (including all retries).

Example: With `StartToCloseTimeout: 5s` and `ScheduleToCloseTimeout: 30s`, each attempt can take up to 5 seconds, but the total time including retries must not exceed 30 seconds.

### Struct Receiver Activities

When an activity is a method on a struct (e.g., for dependency injection), register the struct instance:

```go
// Definition
type PokedexClient struct{}

func (c *PokedexClient) RegisterInPokedexActivity(ctx context.Context, p pokemon.Pokemon) error {
    // ...
}

// Registration in worker
w.RegisterActivity(&PokedexClient{})
```

Temporal will automatically discover all exported methods with valid activity signatures on the struct.

### Getting the Attempt Number

Inside an activity, you can check which attempt you're on:

```go
info := activity.GetInfo(ctx)
attempt := info.Attempt // 1 for first attempt, 2 for first retry, etc.
```

## What to implement

### Activities (`activities.go`)

Implement the same activities as Exercise 3, plus:

- **`ThrowPokeballActivity`** — simulates unreliability:
  - Attempt < 2: return error `"pokeball missed"` (retryable)
  - Attempt >= 3 and capture would fail: return `temporal.NewNonRetryableApplicationError("pokemon fled", "PokemonFled", nil)`
  - Otherwise: return `CaptureResult` based on HP ratio probability

- **`PokedexClient.RegisterInPokedexActivity`** — simulates unreliable API:
  - Attempt < 3: return retryable error
  - Attempt >= 3: succeed

### Workflow (`workflow.go`)

Implement `CapturePokemonWorkflow(ctx, trainerName)`:

1. Default activity options: `StartToCloseTimeout` 10s
2. Call encounter, dodge check, choose, weaken as before
3. For `ThrowPokeballActivity`: use `RetryPolicy{MaximumAttempts: 3, InitialInterval: 1s, BackoffCoefficient: 2.0}`
4. If throw fails (non-retryable error), return `CaptureResult{Success: false, Pokemon: weakened}`
5. If capture succeeds, call `RegisterInPokedexActivity` with:
   - `StartToCloseTimeout: 5s`
   - `ScheduleToCloseTimeout: 30s`
   - `RetryPolicy{MaximumAttempts: 5, InitialInterval: 1s, BackoffCoefficient: 2.0}`
6. Return the `CaptureResult`

## How to test

```bash
go test ./exercises/ex04_errors/...
```

## How to run

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex04_errors/worker/
   ```
3. Start a workflow execution:
   ```bash
   temporal workflow start \
     --type CapturePokemonWorkflow \
     --task-queue pokemon \
     --input '"Ash"'
   ```
4. Check the result in the Temporal UI at http://localhost:8233
