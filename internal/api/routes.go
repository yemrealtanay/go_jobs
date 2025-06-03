package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	r.Post("/jobs", HandleCreateJob)
	r.Get("/jobs", HandleGetJobs)
	r.Get("/jobs/{id}", HandleGetJobById)

	return r
}
