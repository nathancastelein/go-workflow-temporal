# Exercise 9 - Pokemon Journey (Queries)

In this exercise, you will implement a workflow that uses **queries** to expose the current state of a running workflow. The workflow simulates a trainer's Pokemon journey, encountering and capturing Pokemon while allowing external clients to query the journey's progress.

## Concepts

### Queries

Queries allow you to inspect the current state of a running workflow without affecting its execution. A query handler is registered inside the workflow and returns data synchronously.

```go
// Register a query handler inside the workflow
err := workflow.SetQueryHandler(ctx, "query-name", func() (MyType, error) {
    return currentState, nil
})
if err != nil {
    return err
}
```

Queries are read-only: they must not modify workflow state or have side effects. They can be called from an external client even while the workflow is running.

## Your Task

### activities.go

Implement two activities:

- `EncounterPokemonActivity`: picks a random Pokemon from `pokemon.AllPokemon` using `rand.Intn`
- `AttemptCaptureActivity`: attempts to capture a Pokemon; capture succeeds if `p.HP < 100`

### workflow.go

Implement `JourneyWorkflow` that:

1. Sets up activity options with a `StartToCloseTimeout` of 10 seconds
2. Initializes a `JourneyProgress` struct with the trainer name, empty captured list, and status `"exploring"`
3. Registers a query handler named `"progress"` that returns the current `JourneyProgress`
4. Loops 3 times, each iteration:
   - Sets status to `"encountering"`, calls `EncounterPokemonActivity`
   - Increments the encounter count
   - Sets status to `"capturing"`, calls `AttemptCaptureActivity`
   - If capture succeeded, appends the Pokemon to `CapturedPokemon`
   - Sets status back to `"exploring"`
5. Sets status to `"completed"` and returns the final progress

## Validate

Run from the project root:

```bash
go test -v ./exercises/ex09_queries/...
```
