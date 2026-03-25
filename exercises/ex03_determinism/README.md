# Exercise 3 - It Fled! (Determinism)

## Concepts

In this exercise, you will learn about **workflow determinism** - one of the most important rules in Temporal.

- **Determinism constraint**: Workflow code must produce the same sequence of commands every time it is replayed. Temporal replays workflows from their event history for recovery, so any non-deterministic operation (random numbers, current time, UUIDs) will cause the replay to diverge.
- **Two escape hatches**: `workflow.SideEffect` for simple values that just need to be recorded, and **activities** for operations that need full visibility in the UI.

## The Problem

Open `workflow.go` - the workflow from Exercise 2 has been extended with a **flee check**: there's a 30% chance the wild Pokemon flees before you can catch it.

But there's a bug! Can you find it?

**Why this breaks**: When Temporal replays the workflow (e.g., after a worker restart), it re-executes the workflow code from the beginning. The `rand.Float64()` call will return a **different value** on replay, which means the workflow might take a different code path than it originally did. This causes a **non-determinism error** and the workflow gets stuck.

## Step 1: Fix with `workflow.SideEffect`

The quickest way to fix a non-deterministic value in a workflow is `workflow.SideEffect`. It records the result of a function in the event history so that on replay, the recorded value is used instead of re-executing the function.

### What to do

In `workflow.go`, replace the direct `rand.Float64()` call with `workflow.SideEffect`.

Check the documentation of SideEffect here: https://docs.temporal.io/develop/go/side-effects

The rest of the workflow logic stays the same.

### Try it out

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex03_determinism/worker/
   ```
3. Start a workflow execution:
   ```bash
   temporal workflow start \
     --type CapturePokemonWorkflow \
     --task-queue pokemon \
     --input '"Ash"'
   ```
4. Open the Temporal UI at http://localhost:8233 and look at the **event history** of your workflow execution.

Notice how the SideEffect result appears as a **Marker** in the event history. It's recorded, but it doesn't have the same visibility as an activity: there's no clear label, no input/output shown, no duration tracking.

## Step 2: Replace with an Activity

While `workflow.SideEffect` works, activities provide much better **observability** in the Temporal UI: you can see the activity name, its input, its output, its duration, and retry attempts.

### What to do

#### Part A: Implement `FleeCheckActivity` (`activities.go`)

The activity stub is already there. Implement it to return `true` ~30% of the time:

```go
return rand.Float64() < 0.3, nil
```

#### Part B: Update the workflow (`workflow.go`)

Replace the `workflow.SideEffect` call with a call to `FleeCheckActivity`:

```go
var fled bool
err = workflow.ExecuteActivity(ctx, FleeCheckActivity, wildPokemon).Get(ctx, &fled)
```

The fixed workflow should:

1. Encounter a wild Pokemon
2. Call `FleeCheckActivity` - if the Pokemon fled, return `CaptureResult{Success: false, Pokemon: wild}`
3. If not fled: call `ChoosePokemonActivity`, `WeakenActivity`, `ThrowPokeballActivity`
4. Return the `CaptureResult`

### Try it out

1. Restart the worker, then start a new workflow execution
2. Open the Temporal UI and compare the event history with the previous execution (step 1)

Notice how the flee check now appears as a **proper activity** in the event history: you can see `FleeCheckActivity` by name, its input (the Pokemon), and its result (`true` or `false`). Much better for debugging and monitoring!

## How to test

```bash
go test ./exercises/ex03_determinism/...
```

The tests mock `FleeCheckActivity` - they validate the **step 2** solution (activity-based). They will pass once you've completed both steps.

## Takeaway

Both `workflow.SideEffect` and activities fix the determinism issue, but they serve different purposes:

| | `workflow.SideEffect` | Activity |
|---|---|---|
| Determinism | ✅ Recorded in history | ✅ Recorded in history |
| UI visibility | ❌ Opaque marker | ✅ Named, with input/output |
| Retries | ❌ No retry mechanism | ✅ Configurable retries |
| Timeouts | ❌ No timeout | ✅ Configurable timeouts |
| Best for | Simple values (UUID, random) | Logic with I/O or that needs observability |
