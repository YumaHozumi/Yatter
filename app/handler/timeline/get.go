package timeline

import (
	"encoding/json"
	"net/http"
)

func (h *handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entity, err := h.tr.GetPublic(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if entity == nil {
		http.Error(w, "Statuses not found", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
