package job

import (
	"fmt"
	"sync"
	"time"
)

type JobStatus string

const (
	StatusQueued  JobStatus = "queued"
	StatusRunning JobStatus = "running"
	StatusFailed  JobStatus = "failed"
	StatusPending JobStatus = "pending"
	StatusDone    JobStatus = "done"
)

type Job struct {
	JOB_ID      int64                  `json:"job_id"`
	Type        string                 `json:"type"`
	Payload     map[string]interface{} `json:"payload"`
	Status      JobStatus              `json:"status"`
	ProcessedAt *time.Time             `json:"processed_at,omitempty"`
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
	j.Status = StatusQueued
	jobQueue = append(jobQueue, j)

	fmt.Printf("New Job Added: %+v\n", j)
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

func UpdateJobStatus(id int64, status JobStatus) {
	mu.Lock()
	defer mu.Unlock()

	for i, j := range jobQueue {
		if j.JOB_ID == id {
			jobQueue[i].Status = status
			break
		}
	}
}

func MarkJobProcessed(id int64) {
	mu.Lock()
	defer mu.Unlock()
	for i := range jobQueue {
		if jobQueue[i].JOB_ID == id {
			now := time.Now()
			jobQueue[i].ProcessedAt = &now
			break
		}
	}
}
