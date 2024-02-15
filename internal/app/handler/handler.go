package handler

import (
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/alexandremot/go-htmlx/internal/app/service"

	bf "github.com/russross/blackfriday/v2"
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
	case "/readme":
		h.ServeMarkdownPage(w, r, "../../public/README.md")
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
	_, _ = w.Write([]byte(updatedContent))
}

func (h *Handler) ServeMarkdownPage(w http.ResponseWriter, r *http.Request, markdownFilePath string) {
	// Read the Markdown file
	markdownBytes, err := ioutil.ReadFile(markdownFilePath)
	if err != nil {
		http.Error(w, "Markdown file not found", http.StatusNotFound)
		return
	}

	// Convert Markdown to HTML
	extensions := bf.CommonExtensions | bf.AutoHeadingIDs
	renderer := bf.NewHTMLRenderer(bf.HTMLRendererParameters{})
	output := bf.Run(markdownBytes, bf.WithExtensions(extensions), bf.WithRenderer(renderer))

	// set Content-Type to load styles.css
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<!DOCTYPE html>"))
	w.Write([]byte("<html><head><link rel=\"stylesheet\" href=\"static/styles.css\"></head><body>"))
	w.Write(output)
	w.Write([]byte("</body></html>"))
}
