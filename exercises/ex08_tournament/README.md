# Exercise 8 — Pokemon League Tournament

## Concepts

In this exercise, you will learn about **child workflows** — workflows that are started by and managed by a parent workflow. This is the key to composing complex orchestrations from simpler building blocks.

### Child Workflows

Launch a child workflow from a parent using `workflow.ExecuteChildWorkflow`:

```go
childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
    WorkflowExecutionTimeout: 1 * time.Minute,
    ParentClosePolicy:        enumspb.PARENT_CLOSE_POLICY_ABANDON,
})
future := workflow.ExecuteChildWorkflow(childCtx, ChildWorkflow, arg1, arg2)
```

### Futures and Parallel Execution

`ExecuteChildWorkflow` returns a `workflow.ChildWorkflowFuture` immediately — it does not block. This lets you launch multiple child workflows in parallel:

```go
var futures []workflow.ChildWorkflowFuture
for i := 0; i < 4; i++ {
    future := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, trainers[i*2], trainers[i*2+1])
    futures = append(futures, future)
}

// Wait for all results
for _, future := range futures {
    var result pokemon.BattleResult
    err := future.Get(ctx, &result)
    // handle result or error
}
```

### ChildWorkflowOptions

- **`WorkflowExecutionTimeout`**: Maximum time the child workflow can run
- **`ParentClosePolicy`**: What happens to the child if the parent closes (ABANDON, TERMINATE, or REQUEST_CANCEL)

### Error Handling

When `future.Get()` returns an error, the child workflow failed. You can handle this gracefully — for example, advancing the opponent by forfeit in a tournament bracket.

## What to implement

### Activities (`activities.go`)

- `ChoosePokemonActivity`: Look up the trainer's Pokemon in `pokemon.TrainerTeams`
- `WeakenActivity`: Reduce target HP by `attacker.HP / 3`, clamp minimum to 1

### Battle Workflow (`battle.go`)

Implement `BattleWorkflow(ctx, trainer1, trainer2)`:
1. Each trainer chooses their Pokemon (2 activity calls)
2. Pokemon1 attacks Pokemon2 (WeakenActivity)
3. Pokemon2 attacks Pokemon1 (WeakenActivity)
4. Compare remaining HP — higher HP wins

### Tournament Workflow (`tournament.go`)

Implement `TournamentWorkflow(ctx, trainers [8]string)`:
1. **Quarter-finals**: 4 parallel child BattleWorkflow executions
2. **Semi-finals**: 2 parallel child BattleWorkflow executions with quarter-final winners
3. **Final**: 1 child BattleWorkflow execution with semi-final winners
4. Track the bracket (winners of each round) and return the champion

Handle child workflow failures by advancing the opponent by forfeit.

## How to test

```bash
go test ./exercises/ex08_tournament/...
```

## How to run

1. Make sure Temporal Server is running (e.g., `temporal server start-dev`)
2. Start the worker:
   ```bash
   go run ./exercises/ex08_tournament/worker/
   ```
