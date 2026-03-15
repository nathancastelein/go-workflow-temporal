# Exercise 9 - Pokemon League Tournament

In this exercise, you will learn about **child workflows** — workflows started and managed by a parent workflow.

## Concepts

### Child Workflows

Launch a child workflow from a parent using `workflow.ExecuteChildWorkflow`:

```go
childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
    WorkflowExecutionTimeout: 1 * time.Minute,
    ParentClosePolicy:        enumspb.PARENT_CLOSE_POLICY_ABANDON,
})
future := workflow.ExecuteChildWorkflow(childCtx, ChildWorkflow, arg1, arg2)
```

### Parallel Execution with Futures

`ExecuteChildWorkflow` returns a future immediately — it does not block. Launch multiple child workflows in parallel, then collect results:

```go
sf1 := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, "Lorelei", "Bruno")
sf2 := workflow.ExecuteChildWorkflow(childCtx, BattleWorkflow, "Agatha", "Lance")

var result1, result2 pokemon.BattleResult
sf1.Get(ctx, &result1)
sf2.Get(ctx, &result2)
```

## Your Task

### Battle Workflow (`battle.go`)

Implement `BattleWorkflow(ctx, trainer1, trainer2)`:
1. Call `ChoosePokemonActivity` for each trainer
2. Compare HP — higher HP wins
3. Return a `BattleResult{Winner, Loser}`

### Tournament Workflow (`tournament.go`)

Implement `TournamentWorkflow(ctx, lorelei, bruno, agatha, lance)`:
1. **Semi-finals**: 2 parallel `BattleWorkflow` child workflows
2. **Final**: 1 `BattleWorkflow` child workflow with the 2 winners
3. Return a `TournamentResult` with the bracket and champion

## Validate

```bash
go test ./exercises/ex09_tournament/...
```
