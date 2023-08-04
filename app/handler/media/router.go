package media

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	mr repository.Media
}

func NewRouter(mr repository.Media) http.Handler {
	r := chi.NewRouter()

	h := &handler{mr}
	r.Post("/", h.Upload)
	return r
}
