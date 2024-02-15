package handler

import (
	"net/http"
	"path/filepath"

	"github.com/alexandremot/go-htmlx/internal/app/service"
)

type Handler struct {
	Service service.Service // Dependency injection of the service interface
}

func NewHandler(s service.Service) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		// Construct the absolute path to the index.html file
		absPath, err := filepath.Abs("../../public/index.html")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		http.ServeFile(w, r, absPath)
	case "/update":
		h.handleUpdate(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	// Call the service method to get the updated content
	updatedContent, err := h.Service.GetUpdatedContent()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Write the updated content to the response
	_, _ = w.Write([]byte(updatedContent))
}
