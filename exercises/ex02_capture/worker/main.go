package main

import (
	"log"

	ex02 "github.com/nathancastelein/go-workflow-temporal/exercises/ex02_capture"
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

	w.RegisterWorkflow(ex02.CapturePokemonWorkflow)
	w.RegisterActivity(ex02.EncounterWildPokemonActivity)
	w.RegisterActivity(ex02.ChoosePokemonActivity)
	w.RegisterActivity(ex02.WeakenActivity)
	w.RegisterActivity(ex02.ThrowPokeballActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
