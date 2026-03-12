package main

import (
	"log"

	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	ex01 "github.com/nathancastelein/go-workflow-temporal/solutions/ex01_encounter"
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

	w.RegisterWorkflow(ex01.WildEncounterWorkflow)
	w.RegisterActivity(ex01.EncounterWildPokemonActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
