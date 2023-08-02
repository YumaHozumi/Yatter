package accounts

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	//ユーザ検索
	entity, err := h.ar.FindByUsername(r.Context(), username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//取得したentity返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
