package statuses

import (
	"net/http"
	"yatter-backend-go/app/domain/repository"
	"yatter-backend-go/app/handler/auth"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	sr repository.Status
	ar repository.Account
}

func NewRouter(sr repository.Status, ar repository.Account) http.Handler {
	r := chi.NewRouter()

	h := &handler{sr, ar}
	//認証あり
	r.With(auth.Middleware(ar)).Post("/", h.CreateStatus)
	//認証なし
	r.Get("/{id}", h.Find)

	return r
}
