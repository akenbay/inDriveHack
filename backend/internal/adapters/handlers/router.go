package handlers

import (
	"net/http"

	"inDriveHack/internal/services"
)

func NewRouter(inspectorService services.InspectorService) *http.ServeMux {
	mux := http.NewServeMux()
	handler := newHandler(&inspectorService)

	mux.HandleFunc("POST /inspect", handler.upload)

	return mux
}
