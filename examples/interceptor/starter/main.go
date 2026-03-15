package main

import (
	"context"
	"log"
	"log/slog"

	"github.com/nathancastelein/go-workflow-temporal/examples/interceptor"
	"go.temporal.io/sdk/client"
	sdklog "go.temporal.io/sdk/log"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort: "localhost:7233",
		Logger:   sdklog.NewStructuredLogger(slog.Default()),
	})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	we, err := c.ExecuteWorkflow(context.Background(), client.StartWorkflowOptions{
		TaskQueue: "interceptor-example",
	}, interceptor.GreetWorkflow, "Ash")
	if err != nil {
		log.Fatalln("unable to execute workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get workflow result", err)
	}

	slog.Info("workflow completed", "result", result)
}
