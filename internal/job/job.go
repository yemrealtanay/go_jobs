package job

import (
	"fmt"
	"sync"
)

type Job struct {
	JOB_ID  int64                  `json:"job_id"`
	Type    string                 `json:"type"`
	Payload map[string]interface{} `json:"payload"`
}

var (
	jobQueue []Job
	idSeq    int64
	mu       sync.Mutex
)

func AddJob(j Job) Job {
	mu.Lock()
	defer mu.Unlock()
	idSeq++
	j.JOB_ID = idSeq
	jobQueue = append(jobQueue, j)
	fmt.Printf("Added job %d\n", j.JOB_ID)
	return j
}

func GetAllJobs() []Job {
	mu.Lock()
	defer mu.Unlock()
	return jobQueue
}
