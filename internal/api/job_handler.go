package api

import (
	"encoding/json"
	"go_jobs/internal/job"
	"net/http"
)

func HandleCreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var j job.Job
	err := json.NewDecoder(r.Body).Decode(&j)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdJob := job.AddJob(j)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdJob)
}

func HandleGetJobs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	jobs := job.GetAllJobs()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
