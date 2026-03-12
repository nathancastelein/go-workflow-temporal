package main

import (
	"log"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	ex09 "github.com/nathancastelein/go-workflow-temporal/solutions/ex09_queries"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	w := worker.New(c, pokemon.TaskQueue, worker.Options{})

	w.RegisterWorkflow(ex09.JourneyWorkflow)
	w.RegisterActivity(ex09.EncounterPokemonActivity)
	w.RegisterActivity(ex09.AttemptCaptureActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
