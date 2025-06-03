package worker

import (
	"fmt"
	"time"

	"go_jobs/internal/job"
)

var JobQueue chan job.Job

func StartWorkerPool(numWorkers int) {
	JobQueue = make(chan job.Job, numWorkers)

	for i := 1; i <= numWorkers; i++ {
		go worker(i, JobQueue)
	}
}

func worker(id int, jobs <-chan job.Job) {
	for j := range jobs {
		fmt.Printf("[Worker %d] Processing Job #%d - Type: %s\n", id, j.JOB_ID, j.Type)

		// Simule et
		time.Sleep(1 * time.Second)

		fmt.Printf("[Worker %d] Processing Job #%d\n", id, j.JOB_ID)
	}
}
