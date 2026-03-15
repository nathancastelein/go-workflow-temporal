# Exercise 8 - Team Rocket is Watching

In this exercise, you will implement a simple **interceptor** that spies on all activity executions — Team Rocket wants to know every time a trainer catches a Pokemon!

## Context

The workflow and activities are already implemented:
- `CatchPokemonWorkflow` encounters a wild Pokemon and throws a Pokeball
- `EncounterActivity` returns a random wild Pokemon
- `ThrowPokeballActivity` attempts to catch it (30% success rate)

Your job is to implement the **Team Rocket interceptor** that records every activity execution.

Interceptors work like middleware. There are two steps:
1. **Infiltrate**: register yourself in the activity chain so all calls go through you
2. **Spy**: observe each activity execution and record what happened

## Your Task

### interceptors.go

`InterceptActivity` is already implemented — it inserts Team Rocket into the activity chain at startup by storing `next` and returning itself. Without this, activities would run without Team Rocket seeing anything.

**`ExecuteActivity`** — spy on each activity execution:
- Now that you're in the chain, every activity call goes through you
- Delegate to `tr.Next.ExecuteActivity(ctx, in)` to let the activity run
- Get the activity name with `activity.GetInfo(ctx).ActivityType.Name`
- Append a `TeamRocketReport{ActivityName, Success}` where `Success` is `err == nil`
- Return the result and error unchanged

## Validate

```bash
go test ./exercises/ex08_interceptors/...
```
