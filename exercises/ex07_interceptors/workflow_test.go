package ex07_interceptors

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/worker"
)

type InterceptorTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
	env     *testsuite.TestWorkflowEnvironment
	factory *WorkerInterceptorFactory
}

func (s *InterceptorTestSuite) SetupTest() {
	s.factory = NewWorkerInterceptorFactory()
	s.env = s.NewTestWorkflowEnvironment()
	s.env.SetWorkerOptions(worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{s.factory},
	})
	s.env.RegisterActivity(SimpleActivity)
}

func (s *InterceptorTestSuite) AfterTest(suiteName, testName string) {
	s.env.AssertExpectations(s.T())
}

func TestInterceptorSuite(t *testing.T) {
	suite.Run(t, new(InterceptorTestSuite))
}

func (s *InterceptorTestSuite) TestSimpleWorkflow_WithTrainerName() {
	// Arrange
	// (setup done in SetupTest)

	// Act
	s.env.ExecuteWorkflow(SimpleWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	var result string
	s.NoError(s.env.GetWorkflowResult(&result))
	s.Equal("done", result)
}

func (s *InterceptorTestSuite) TestSimpleWorkflow_WithoutTrainerName() {
	// Arrange
	// (setup done in SetupTest)

	// Act
	s.env.ExecuteWorkflow(SimpleWorkflow, "")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	err := s.env.GetWorkflowError()
	s.Error(err)
	s.Contains(err.Error(), "trainer name argument required")
}

func (s *InterceptorTestSuite) TestSpyActivityInterceptor_LogsActivity() {
	// Arrange
	// (setup done in SetupTest)

	// Act
	s.env.ExecuteWorkflow(SimpleWorkflow, "Ash")

	// Assert
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())

	s.Require().Len(s.factory.SpyInterceptor.Reports, 1)
	report := s.factory.SpyInterceptor.Reports[0]
	s.Equal("SimpleActivity", report.ActivityName)
	s.NoError(report.Error)
	s.Greater(report.Duration, 0*time.Millisecond)
}
