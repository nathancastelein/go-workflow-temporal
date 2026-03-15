# Exercise 3 - It Fled! (Determinism)

## Concepts

In this exercise, you will learn about **workflow determinism** - one of the most important rules in Temporal.

- **Determinism constraint**: Workflow code must produce the same sequence of commands every time it is replayed. Temporal replays workflows from their event history for recovery, so any non-deterministic operation (random numbers, current time, UUIDs) will cause the replay to diverge.
- **Activities are the escape hatch**: Since activity results are recorded in the event history, any non-deterministic logic (randomness, I/O, etc.) should live inside activities.

## The Problem

Open `workflow.go` - the workflow from Exercise 2 has been extended with a **flee check**: there's a 30% chance the wild Pokemon flees before you can catch it.

But there's a bug! Can you find it?

**Why this breaks**: When Temporal replays the workflow (e.g., after a worker restart), it re-executes the workflow code from the beginning. The `rand.Float64()` call will return a **different value** on replay, which means the workflow might take a different code path than it originally did. This causes a **non-determinism error** and the workflow gets stuck.

## What to implement

### Part 1: Implement `FleeCheckActivity` (`activities.go`)

The activities from Exercise 2 are already provided.

Create a new activity to safely encapsulate the randomness:

1. **`FleeCheckActivity`** - returns `true` ~30% of the time using `rand.Float64() < 0.3`

### Part 2: Fix the workflow (`workflow.go`)

Replace the non-deterministic `rand.Float64()` call with a call to your new `FleeCheckActivity`.

The fixed workflow should:

1. Encounter a wild Pokemon
2. Call `FleeCheckActivity` - if the Pokemon fled, return `CaptureResult{Success: false, Pokemon: wild}`
3. If not fled: call `ChoosePokemonActivity`, `WeakenActivity`, `ThrowPokeballActivity`
4. Return the `CaptureResult`

## How to test

```bash
go test ./exercises/ex03_determinism/...
```

The tests mock `FleeCheckActivity` - they will fail until you fix the workflow to use the activity instead of `rand.Float64()`.

## How to run

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
4. Check the result in the Temporal UI at http://localhost:8233
