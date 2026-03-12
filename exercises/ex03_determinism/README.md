# Exercise 3 — It Dodged! (Determinism)

## Concepts

In this exercise, you will learn about **workflow determinism** — one of the most important rules in Temporal.

- **Determinism constraint**: Workflow code must produce the same sequence of commands every time it is replayed. Temporal replays workflows from their event history for recovery, so any non-deterministic operation (random numbers, current time, UUIDs) will cause the replay to diverge.
- **Activities are the escape hatch**: Since activity results are recorded in the event history, any non-deterministic logic (randomness, I/O, etc.) should live inside activities.
- **`workflow.SideEffect`**: An alternative for lightweight non-deterministic operations that don't need retry semantics.

## The Trap

Here is code that **looks correct** but **breaks Temporal determinism**:

```go
// DON'T DO THIS — this is non-deterministic!
func BrokenCapturePokemonWorkflow(ctx workflow.Context, trainerName string) (pokemon.CaptureResult, error) {
    // Using math/rand directly in a workflow breaks replay determinism
    dodged := rand.Float64() < 0.3
    if dodged {
        return pokemon.CaptureResult{Success: false}, nil
    }
    // ...
}
```

**Why this breaks**: When Temporal replays the workflow (e.g., after a worker crash), it re-executes the workflow code from the beginning. The `rand.Float64()` call will return a **different value** on replay, which means the workflow might take a different code path than it originally did. This causes a **non-determinism error** and the workflow gets stuck.

**The fix**: Move the randomness into an activity (`DodgeCheckActivity`). Activity results are recorded in the event history, so on replay Temporal uses the recorded result instead of re-executing the activity code.

## What to implement

### Activities (`activities.go`)

Implement 5 activities:

1. **`EncounterWildPokemonActivity`** — picks a random Pokemon from `pokemon.AllPokemon`
2. **`ChoosePokemonActivity`** — looks up the trainer's Pokemon from `pokemon.TrainerTeams`
3. **`WeakenActivity`** — reduces target HP by `attacker.HP / 3`, clamp min 1
4. **`ThrowPokeballActivity`** — capture probability based on HP ratio: `1.0 - (target.HP / target.MaxHP)`
5. **`DodgeCheckActivity`** — returns `true` ~30% of the time using `rand.Float64() < 0.3`

### Workflow (`workflow.go`)

Implement `CapturePokemonWorkflow(ctx, trainerName)`:

1. Set activity options with `StartToCloseTimeout` of 10 seconds
2. Call `EncounterWildPokemonActivity` to encounter a wild Pokemon
3. Call `DodgeCheckActivity` — if the Pokemon dodged, return `CaptureResult{Success: false, Pokemon: wild}`
4. If not dodged: call `ChoosePokemonActivity`, `WeakenActivity`, `ThrowPokeballActivity`
5. Return the `CaptureResult`

## How to test

```bash
go test ./exercises/ex03_determinism/...
```

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
