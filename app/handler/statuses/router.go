package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	sr repository.Status
}

func NewRouter(sr repository.Status) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr}

	r.Post("/", h.CreateStatus)

	return r
}
