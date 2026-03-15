package main

import (
	"log"
	"log/slog"

	example "github.com/nathancastelein/go-workflow-temporal/examples/interceptor"
	"go.temporal.io/sdk/client"
	sdkinterceptor "go.temporal.io/sdk/interceptor"
	sdklog "go.temporal.io/sdk/log"
	"go.temporal.io/sdk/worker"
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

	w := worker.New(c, "interceptor-example", worker.Options{
		Interceptors: []sdkinterceptor.WorkerInterceptor{&example.LoggingWorkerInterceptor{}},
	})

	w.RegisterWorkflow(example.GreetWorkflow)
	w.RegisterActivity(example.GreetActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
