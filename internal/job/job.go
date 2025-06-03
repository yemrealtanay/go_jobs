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

	fmt.Printf("New Job Added - ID: %d\n", j.JOB_ID)
	return j
}

func GetAllJobs() []Job {
	mu.Lock()
	defer mu.Unlock()
	return jobQueue
}

func GetJobById(id int64) (*Job, bool) {
	mu.Lock()
	defer mu.Unlock()

	for _, j := range jobQueue {
		if j.JOB_ID == id {
			return &j, true
		}
	}
	return nil, false
}
