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
- If the Pokemon has an evolution, return the evolved form
- If not, return the same Pokemon unchanged

### workflow.go

Implement `EvolutionWorkflow` that:
1. Creates signal channels for "feed" and "cancel"
2. Creates a timer with the given duration
3. Uses a selector to wait for the first event:
   - **Timer fires**: Call `EvolveActivity`, return with `Trigger: "timer"`
   - **Feed signal**: Call `EvolveActivity`, return with `Trigger: "feed"`
   - **Cancel signal**: Return the Pokemon unchanged with `Trigger: "cancelled"`

## Validate

Run from the project root:

```bash
go test -v ./exercises/ex06_signals/...
```
