# Pokemon Temporal Training Workshop

A hands-on workshop to learn [Temporal](https://temporal.io) with Go, using a Pokemon theme.

## Prerequisites

- **Go 1.22+**: [Install Go](https://go.dev/dl/)
- **Temporal CLI**: [Install Temporal CLI](https://docs.temporal.io/cli#install)

## Getting Started

### 1. Start the Temporal dev server

```bash
temporal server start-dev
```

This starts a local Temporal server with a Web UI at http://localhost:8233.

### 2. Run an exercise

Each exercise is in `exercises/exNN_name/`. Open the `README.md` in the exercise folder for instructions.

### 3. Run tests to validate your implementation

```bash
# Test your implementation
go test ./exercises/ex01_encounter/...

# Check the solution
go test ./solutions/ex01_encounter/...
```

### 4. Run the worker (for exercises with a worker)

```bash
# Start the worker
go run ./exercises/ex01_encounter/worker/

# In another terminal, use the Temporal CLI to start a workflow
temporal workflow start --task-queue pokemon --type WildEncounterWorkflow
```

## Discovery Exercise

Before diving into Temporal, explore two different workflow approaches:

```bash
# Run Approach A
cd discovery/approach-a && go run .

# Run Approach B
cd discovery/approach-b && go run .
```

Compare how each program implements the same capture workflow differently.

## HelloWorld Example

A complete Temporal HelloWorld to run after setting up Temporal:

```bash
# Terminal 1: Start the worker
go run ./examples/helloworld/worker/

# Terminal 2: Run the starter
go run ./examples/helloworld/starter/
```

## Exercises

| # | Name | Concepts |
|---|------|----------|
| 1 | Tall Grass Encounter | Workflow, Activity, Worker |
| 2 | I Choose You! | Multiple activities, Data passing |
| 3 | It Fled! | Determinism |
| 4 | The Pokeball Missed! | Error handling, Retries, Timeouts |
| 5 | Testing Your Pokemon | Temporal test framework |
| 6 | Evolution Chamber | Signals, Timers, Selectors |
| 7 | Pokemon Journey | Queries |
| 8 | Team Rocket is Watching | Interceptors |
| 9 | Pokemon League Tournament | Child workflows, Parallelism |

## Project Structure

```
go-workflow-temporal/
├── pokemon/          # Shared domain types and data (read-only)
├── discovery/        # Discovery exercise: two workflow approaches
│   ├── approach-a/   # Approach A — run with: go run .
│   └── approach-b/   # Approach B — run with: go run .
├── examples/         # Provided examples
│   └── helloworld/   # HelloWorld Temporal workflow
│       ├── activities.go
│       ├── workflow.go
│       ├── worker/   # Start with: go run .
│       └── starter/  # Start with: go run .
├── exercises/        # Exercise stubs — implement these!
│   └── exNN_name/
│       ├── README.md
│       ├── *.go          # Stubbed files to complete
│       ├── *_test.go     # Pre-written tests
│       └── worker/       # Worker main (some exercises)
└── solutions/        # Reference implementations
    └── exNN_name/
```
