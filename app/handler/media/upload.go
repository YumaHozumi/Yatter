package media

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"yatter-backend-go/app/domain/object"
)

func (h *handler) Upload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//今回はローカルにディレクトリ作成してファイルを保存しておく
	dir := "./tmp/"
	os.MkdirAll(dir, os.ModePerm) // Make sure the directory exists

	path := filepath.Join(dir, header.Filename)
	out, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	media := new(object.Media)
	media.MediaURL = path

	ctx := r.Context()

	id, err := h.mr.UploadMedia(ctx, media)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	entity, err := h.mr.FindMedia(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if entity == nil {
		http.Error(w, "media not found", http.StatusBadRequest)
		return
	}

	//取得したentity返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
