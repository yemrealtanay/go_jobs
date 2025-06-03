package main

import (
	"log"
	"net/http"

	"go_jobs/internal/api"
	"go_jobs/internal/worker"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	worker.StartWorkerPool(3)
	r := api.RegisterRoutes()

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
