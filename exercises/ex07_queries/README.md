# Exercise 7 - Pokemon Journey (Queries)

In this exercise, you will add a **query handler** to an existing workflow to expose the current journey progress to external clients.

## Concepts

### Queries

Queries allow you to inspect the current state of a running workflow without affecting its execution. A query handler is registered inside the workflow and returns data synchronously.

```go
err := workflow.SetQueryHandler(ctx, "query-name", func() (MyType, error) {
    return currentState, nil
})
if err != nil {
    return err
}
```

Queries are read-only: they must not modify workflow state or have side effects. They can be called from an external client even while the workflow is running.

## Context

The workflow encounters a wild Pokemon and attempts to capture it. The encounter and capture logic is already implemented — the workflow returns a `pokemon.CaptureResult`.

Your goal is to **track the journey progress** and expose it via a query handler so external clients can monitor the workflow in real-time.

## Your Task

### workflow.go — Track progress and register a query handler

A `JourneyProgress` struct is already defined in `types.go`. Look for the `TODO` comments in the workflow:

1. Create a `JourneyProgress` variable with the trainer name and initial status `"exploring"`
2. Register a query handler named `"progress"` that returns the current `JourneyProgress`
3. Update the progress status before each activity (`"encountering"`, `"capturing"`)
4. After the capture attempt, set the final status (`"captured"` or `"escaped"`) and store the encountered Pokemon

## Validate

```bash
go test ./exercises/ex07_queries/...
```

## Try it with a real Temporal server

Once your worker is running against a local Temporal server, you can interact with the workflow using the CLI.

Start the workflow (with a long timer so you have time to send signals):

```bash
temporal workflow start \
  --task-queue pokemon \
  --type JourneyWorkflow \
  --input '"Sacha"'
```

### Query the result

```bash
temporal workflow query \
  --workflow-id <WORKFLOW_ID> \
  --type progress
```