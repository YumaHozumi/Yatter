package statuses

import (
	"encoding/json"
	"net/http"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
)

type AddRequest struct {
	Status string `json:"status"`
}

type Response struct {
	Id         int64          `json:"id"`
	Account    object.Account `json:"account"`
	Content    string         `json:"content"`
	Created_at time.Time      `json:"created_at"`
}

func (h *handler) CreateStatus(w http.ResponseWriter, r *http.Request) {
	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status := new(object.Status)

	status.Content = req.Status

	account := auth.AccountOf(r)

	ctx := r.Context()

	// save the new status
	id, err := h.sr.CreateStatus(ctx, status, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity, err := h.sr.FindStatusByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var res Response
	res.Id = entity.ID
	res.Account = *account
	res.Content = entity.Content
	res.Created_at = entity.CreateAt

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
