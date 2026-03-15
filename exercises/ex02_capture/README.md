# Exercise 2 - I Choose You!

## Concepts

In this exercise, you will learn how to:

- Chain **multiple activities** in a workflow
- Pass **data between activities** - the output of one activity becomes the input of the next
- Configure **activity options** that apply to all activities in a workflow

## What to implement

### Activities (`activities.go`)

`EncounterWildPokemonActivity` is already provided from Exercise 1.

Implement 3 activities:

1. **`ChoosePokemonActivity(trainerName)`** - looks up the trainer's Pokemon in `pokemon.TrainerTeams`. Returns an error if the trainer is unknown.
2. **`WeakenActivity(attacker, target)`** - reduces `target.HP` by `attacker.HP / 3`. Clamp HP to a minimum of 1.
3. **`ThrowPokeballActivity(target)`** - capture probability is `1.0 - (target.HP / target.MaxHP)`. Returns a `pokemon.CaptureResult`.

### Workflow (`workflow.go`)

Implement `CapturePokemonWorkflow(trainerName)` - it should:

1. Create activity options with a `StartToCloseTimeout` of 10 seconds
2. Encounter a wild Pokemon
3. Choose the trainer's Pokemon
4. Weaken the wild Pokemon using the trainer's Pokemon
5. Throw a Pokeball at the weakened Pokemon
6. Return the `CaptureResult`

## How to test

```bash
go test ./exercises/ex02_capture/...
```

## How to run

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex02_capture/worker/
   ```
3. Start a workflow execution:
   ```bash
   temporal workflow start \
     --type CapturePokemonWorkflow \
     --task-queue pokemon \
     --input '"Ash"'
   ```
4. Check the result in the Temporal UI at http://localhost:8233
