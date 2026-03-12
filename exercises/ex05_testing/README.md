# Exercise 5 - Testing Your Pokemon

In this exercise, you will write tests for a Temporal workflow using the Temporal test framework. The workflow and activity code is already provided -- your job is to write the tests!

## The Temporal Test Framework

Temporal provides a test framework built on top of `testify/suite` that lets you test workflows without running a Temporal server.

### Key Components

**WorkflowTestSuite**: The base test suite that provides a `TestWorkflowEnvironment`.

```go
type WorkflowTestSuite struct {
    suite.Suite
    testsuite.WorkflowTestSuite
    env *testsuite.TestWorkflowEnvironment
}
```

**TestWorkflowEnvironment**: A simulated Temporal environment where you can:
- Mock activities
- Execute workflows
- Assert results

### Mocking Activities

Use `env.OnActivity()` to mock activity behavior:

```go
s.env.OnActivity(MyActivity, mock.Anything, arg1, arg2).Return(result, nil)
```

- `mock.Anything` is used for the `context.Context` parameter (first argument)
- Following arguments should match what the workflow passes
- `.Return(value, error)` sets the mock return values

### Verifying Call Counts

Use `.Times(N)` to assert an activity is called exactly N times:

```go
s.env.OnActivity(MyActivity, mock.Anything).Return(result, nil).Times(1)
```

The `AfterTest` method calls `env.AssertExpectations(s.T())` which verifies all mock expectations were met.

### Executing and Checking Workflows

```go
// Execute the workflow
s.env.ExecuteWorkflow(MyWorkflow, arg1, arg2)

// Check completion
s.True(s.env.IsWorkflowCompleted())
s.NoError(s.env.GetWorkflowError())

// Get the result
var result MyResultType
s.NoError(s.env.GetWorkflowResult(&result))
```

### Direct Activity Tests

You can also test activities directly without the test environment:

```go
result, err := MyActivity(nil, arg1)
s.NoError(err)
s.NotEmpty(result.Name)
```

## Your Task

Open `workflow_test.go` and implement the five test methods. Each has TODO comments with hints.

## Validate

Run from the project root:

```bash
bash exercises/ex05_testing/validate.sh
```
