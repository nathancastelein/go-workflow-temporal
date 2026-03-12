package main

import (
	"log"

	ex07 "github.com/nathancastelein/go-workflow-temporal/exercises/ex07_interceptors"
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

	factory := ex07.NewWorkerInterceptorFactory()

	w := worker.New(c, pokemon.TaskQueue, worker.Options{
		Interceptors: []interceptor.WorkerInterceptor{factory},
	})

	w.RegisterWorkflow(ex07.SimpleWorkflow)
	w.RegisterActivity(ex07.SimpleActivity)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Unable to start worker: %v", err)
	}
}
