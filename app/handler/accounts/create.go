package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//panic("Must Implement Account Registration")

	ctx := r.Context()

	//同じユーザネームが重複して登録されようとしたら弾く
	if entity, err := h.ar.FindByUsername(ctx, account.Username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if entity != nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	//ユーザ作成する
	if err := h.ar.CreateUser(ctx, account); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//dbからユーザ取得
	entity, err := h.ar.FindByUsername(ctx, account.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
