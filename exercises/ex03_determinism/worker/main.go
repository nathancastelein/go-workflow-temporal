package main

import (
	"log"

	ex03 "github.com/nathancastelein/go-workflow-temporal/exercises/ex03_determinism"
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
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

	w.RegisterWorkflow(ex03.CapturePokemonWorkflow)
	w.RegisterActivity(ex03.EncounterWildPokemonActivity)
	w.RegisterActivity(ex03.ChoosePokemonActivity)
	w.RegisterActivity(ex03.WeakenActivity)
	w.RegisterActivity(ex03.ThrowPokeballActivity)
	w.RegisterActivity(ex03.DodgeCheckActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
