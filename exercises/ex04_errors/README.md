# Exercise 4 - The Pokeball Missed! (Error Handling)

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

Sometimes an error is permanent - retrying won't help. Use `temporal.NewNonRetryableApplicationError` to stop retries immediately:

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

The activities from Exercise 3 are already provided (`EncounterWildPokemonActivity`, `ChoosePokemonActivity`, `WeakenActivity`, `FleeCheckActivity`). Note that `ThrowPokeballActivity` has been replaced with a new version for this exercise.

Implement 2 new activities:

- **`ThrowPokeballActivity`** - simulates unreliability using the attempt number:
  - Attempt < 2: return a retryable error ("pokeball missed")
  - Attempt >= 3 and capture would fail: return a non-retryable error (the pokemon fled)
  - Otherwise: return a `CaptureResult` based on HP ratio probability

- **`PokedexClient.RegisterInPokedexActivity`** - simulates an unreliable API:
  - Attempt < 3: return a retryable error
  - Attempt >= 3: succeed

### Workflow (`workflow.go`)

The base capture flow from Exercise 3 is already pre-filled (encounter, flee check, choose, weaken, throw). You only need to add error handling:

1. For `ThrowPokeballActivity` (Step 5): create separate activity options with a `RetryPolicy{MaximumAttempts: 3, InitialInterval: 1s, BackoffCoefficient: 2.0}`
2. If throw fails (non-retryable error = pokemon fled), return `CaptureResult{Success: false, Pokemon: weakened}`
3. If capture succeeds, call `RegisterInPokedexActivity` with its own activity options:
   - `StartToCloseTimeout: 5s`
   - `ScheduleToCloseTimeout: 30s`
   - `RetryPolicy{MaximumAttempts: 5, InitialInterval: 1s, BackoffCoefficient: 2.0}`

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
