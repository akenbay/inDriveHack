package handlers

import (
	"inDriveHack/internal/services"
	"net/http"
)

type Handler struct {
	inspector *services.InspectorService
}

func newHandler(inspector *services.InspectorService) *Handler {
	return &Handler{inspector: inspector}
}

func (h *Handler) upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		respondError(w, "Failed to parse file, the size is too big.", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["images"]
	if len(files) == 0 {
		respondError(w, "No images were provided.", http.StatusBadRequest)
		return
	}

	result, err := h.inspector.Inspect(files[0])

	if err != nil {
		respondError(w, "Error when analyzing the image.", http.StatusBadRequest)
		return
	}

	respondJSON(w, result, http.StatusOK)

}
