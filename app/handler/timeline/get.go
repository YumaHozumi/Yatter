package timeline

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *handler) GetPublic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	//URLからクエリ取得
	maxIDParam := r.URL.Query().Get("max_id")
	sinceIDParam := r.URL.Query().Get("since_id")
	limitParam := r.URL.Query().Get("limit")

	// Convert max_id and since_id from string to int64
	var maxID, sinceID int64
	if maxIDParam != "" {
		maxID, _ = strconv.ParseInt(maxIDParam, 10, 64)
	}
	if sinceIDParam != "" {
		sinceID, _ = strconv.ParseInt(sinceIDParam, 10, 64)
	}

	// Set limit depending on query parameter
	limit := 40
	if limitParam != "" {
		customLimit, _ := strconv.Atoi(limitParam)
		if customLimit > 0 && customLimit <= 80 {
			limit = customLimit
		}
	}

	entity, err := h.tr.GetPublic(ctx, limit, maxID, sinceID)
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
