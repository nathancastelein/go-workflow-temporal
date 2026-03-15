package main

import (
	"log"

	ex08 "github.com/nathancastelein/go-workflow-temporal/exercises/ex08_interceptors"
	"github.com/nathancastelein/go-workflow-temporal/pokemon"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/interceptor"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalf("Unable to create Temporal client: %v", err)
	}
	defer c.Close()

	teamRocket := ex08.NewTeamRocketInterceptor()

	w := worker.New(c, pokemon.TaskQueue, worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{teamRocket},
	})

	w.RegisterWorkflow(ex08.CatchPokemonWorkflow)
	w.RegisterActivity(ex08.EncounterActivity)
	w.RegisterActivity(ex08.ThrowPokeballActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
