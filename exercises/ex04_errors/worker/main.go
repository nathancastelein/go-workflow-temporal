package main

import (
	"log"

	ex04 "github.com/nathancastelein/go-workflow-temporal/exercises/ex04_errors"
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

	w.RegisterWorkflow(ex04.CapturePokemonWorkflow)
	w.RegisterActivity(ex04.EncounterWildPokemonActivity)
	w.RegisterActivity(ex04.ChoosePokemonActivity)
	w.RegisterActivity(ex04.WeakenActivity)
	w.RegisterActivity(ex04.DodgeCheckActivity)
	w.RegisterActivity(ex04.ThrowPokeballActivity)
	w.RegisterActivity(&ex04.PokedexClient{})

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
