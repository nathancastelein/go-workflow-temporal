package ex07_interceptors

import (
	"context"
)

// SimpleActivity is a basic activity used for testing interceptors.
func SimpleActivity(ctx context.Context) (string, error) {
	return "done", nil
}
