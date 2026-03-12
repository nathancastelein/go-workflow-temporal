package main

import (
	"log"
	"log/slog"

	"github.com/nathancastelein/go-workflow-temporal/examples/helloworld"
	"go.temporal.io/sdk/client"
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

	w := worker.New(c, "helloworld", worker.Options{})

	w.RegisterWorkflow(helloworld.Helloworld)
	w.RegisterActivity(helloworld.SayHelloToTrainer)
	w.RegisterActivity(helloworld.SayHelloToPokemon)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start worker", err)
	}
}
