package main

import (
	"log"

	ex08 "github.com/nathancastelein/go-workflow-temporal/exercises/ex08_tournament"
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

	w.RegisterWorkflow(ex08.TournamentWorkflow)
	w.RegisterWorkflow(ex08.BattleWorkflow)
	w.RegisterActivity(ex08.ChoosePokemonActivity)
	w.RegisterActivity(ex08.WeakenActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
