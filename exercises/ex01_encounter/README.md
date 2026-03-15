# Exercise 1 - Tall Grass Encounter

## Concepts

In this exercise, you will learn the three fundamental building blocks of Temporal:

- **Activity**: A function that performs a single unit of work (e.g., calling an API, reading a database). Activities can fail and be retried automatically by Temporal.
- **Workflow**: A function that orchestrates activities. Workflows are durable - if the process crashes, Temporal replays the workflow from where it left off.
- **Worker**: A process that polls Temporal for work, executing workflows and activities.

## What to implement

### Activity (`activities.go`)

Implement `EncounterWildPokemonActivity` - it should pick a random Pokemon from `pokemon.AllPokemon` and return it.

### Workflow (`workflow.go`)

Implement `WildEncounterWorkflow` - it should:

1. Create activity options with a `StartToCloseTimeout` of 10 seconds
2. Call `EncounterWildPokemonActivity`
3. Return the encountered Pokemon

## How to test

```bash
go test ./exercises/ex01_encounter/...
```

## How to run

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex01_encounter/worker/
   ```
3. Start a workflow execution:
   ```bash
   temporal workflow start \
     --type WildEncounterWorkflow \
     --task-queue pokemon
   ```
4. Check the result in the Temporal UI at http://localhost:8233
