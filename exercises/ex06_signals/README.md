# Exercise 6 - Evolution Chamber

In this exercise, you will implement a workflow that uses **signals** and **timers** with the **selector pattern** to handle multiple asynchronous events.

## Concepts

### Signals

Signals allow you to send data to a running workflow from the outside. A workflow can listen for signals using signal channels.

```go
// Create a signal channel
ch := workflow.GetSignalChannel(ctx, "signal-name")

// Receive from the channel (blocks until a signal is received)
var value MyType
ch.Receive(ctx, &value)
```

### Timers

Timers create a future that resolves after a specified duration.

```go
timerFuture := workflow.NewTimer(ctx, 5*time.Second)

// Wait for the timer to fire
timerFuture.Get(ctx, nil)
```

### Selectors

Selectors let you wait for **the first** of multiple events (signals, timers, activities) to complete. This is similar to a `select` statement in Go, but for Temporal futures and channels.

```go
sel := workflow.NewSelector(ctx)

// Wait for a timer
sel.AddFuture(timerFuture, func(f workflow.Future) {
    // Timer fired!
})

// Wait for a signal
sel.AddReceive(signalCh, func(ch workflow.ReceiveChannel, more bool) {
    ch.Receive(ctx, nil)
    // Signal received!
})

// Block until the first event fires
sel.Select(ctx)
```

## Your Task

### activities.go

Implement `EvolveActivity` that looks up a Pokemon's evolution in `pokemon.EvolutionMap`:
- If the Pokemon has an evolution, return the **evolved** `pokemon.Pokemon`
- If not, return the **same** Pokemon unchanged

Note: the activity only handles the evolution lookup. It returns a `pokemon.Pokemon`, not an `EvolutionResult` — the workflow is responsible for wrapping it into an `EvolutionResult` with the right `Trigger` and `Evolved` fields.

### workflow.go

Implement `EvolutionWorkflow` that:
1. Creates two signal channels: `"feed"` and `"cancel"`
2. Creates a timer with the given duration
3. Uses a selector to wait for the first event:
   - **Timer fires**: Call `EvolveActivity` with the Pokemon, wrap the returned Pokemon into an `EvolutionResult{Evolved: true, Trigger: "timer"}`
   - **Feed signal**: Call `EvolveActivity` with the Pokemon, wrap the returned Pokemon into an `EvolutionResult{Evolved: true, Trigger: "feed"}`
   - **Cancel signal**: Return the Pokemon unchanged as `EvolutionResult{Evolved: false, Trigger: "cancelled"}`

If `EvolveActivity` returns an error, return `Evolved: false` with the original Pokemon.

## Validate

Run from the project root:

```bash
go test -v ./exercises/ex06_signals/...
```

## Try it with a real Temporal server

Once your worker is running against a local Temporal server, you can interact with the workflow using the CLI.

Start the workflow (with a long timer so you have time to send signals):

```bash
temporal workflow start \
  --task-queue pokemon \
  --type EvolutionWorkflow \
  --input '{"Name":"Charmander","Type":"Fire","HP":39,"MaxHP":39}' \
  --input '120000000000'
```

Send a **feed** signal to trigger evolution immediately:

```bash
temporal workflow signal --workflow-id <workflow-id> --name feed
```

Or send a **cancel** signal to abort:

```bash
temporal workflow signal --workflow-id <workflow-id> --name cancel
```

Check the result:

```bash
temporal workflow show --workflow-id <workflow-id>
```
