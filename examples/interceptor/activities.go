package interceptor

import (
	"context"
	"fmt"
)

func GreetActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s!", name), nil
}
