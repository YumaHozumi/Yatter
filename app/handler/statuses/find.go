package statuses

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) Find(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	entity, err := h.sr.FindStatusByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if entity == nil {
		http.Error(w, "Status not found", http.StatusBadRequest)
		return
	}

	//statusのアカウントIDからアカウント取得
	account, err := h.ar.FindByUserID(ctx, entity.AccountID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if entity == nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	var res Response
	res.Id = entity.ID
	res.Account = *account
	res.Content = entity.Content
	res.Created_at = entity.CreateAt

	//取得したentity返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
