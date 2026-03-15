# Exercise 5 - Testing Pokemon Evolution

In this exercise, you will write tests for a Temporal workflow using the Temporal test framework. The workflow and activity code is already provided -- your job is to write the tests!

## The Temporal Test Framework

Temporal provides a test framework that lets you test workflows without running a Temporal server.

### Key Components

**TestWorkflowEnvironment**: A simulated Temporal environment where you can mock activities, execute workflows, and assert results.

```go
testSuite := &testsuite.WorkflowTestSuite{}
env := testSuite.NewTestWorkflowEnvironment()
```

### Mocking Activities

Use `env.OnActivity()` to mock activity behavior:

```go
env.OnActivity(MyActivity, mock.Anything, arg1, arg2).Return(result, nil)
```

- `mock.Anything` is used for the `context.Context` parameter (first argument)
- Following arguments should match what the workflow passes
- `.Return(value, error)` sets the mock return values

### Executing and Checking Workflows

```go
// Execute the workflow
env.ExecuteWorkflow(MyWorkflow, arg1, arg2)

// Check completion
require.True(t, env.IsWorkflowCompleted())
require.NoError(t, env.GetWorkflowError())

// Get the result
var result MyResultType
require.NoError(t, env.GetWorkflowResult(&result))
```

### Testing Activities in Isolation

To test activities without a workflow, use `NewTestActivityEnvironment`:

```go
testSuite := &testsuite.WorkflowTestSuite{}
activityEnv := testSuite.NewTestActivityEnvironment()
activityEnv.RegisterActivity(MyActivity)

// Execute the activity
encodedResult, err := activityEnv.ExecuteActivity(MyActivity, arg1, arg2)

// Decode and assert the result
require.NoError(t, err)
var result MyResultType
require.NoError(t, encodedResult.Get(&result))
```

## Your Task

The `EvolvePokemonWorkflow` orchestrates 3 activities:
1. `FetchPokemonActivity` - fetches a Pokemon by name
2. `CheckEvolutionActivity` - checks if the Pokemon can evolve
3. `EvolvePokemonActivity` - performs the evolution

Write the following tests:

- **`workflow_test.go`**: `TestWorkflow_SuccessfulEvolution` - mock all 3 activities for a successful Pikachu → Raichu evolution
- **`activities_test.go`**: `TestFetchPokemonActivity_KnownPokemon` - test with a known name, assert the correct Pokemon is returned
- **`activities_test.go`**: `TestFetchPokemonActivity_UnknownPokemon` - test with an unknown name, assert an error is returned

## Validate

Run from the exercise directory:

```bash
go test .
```
