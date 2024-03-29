package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/alexandremot/go-htmlx/infrastructure/web"
	"github.com/alexandremot/go-htmlx/internal/app/handler"
	"github.com/alexandremot/go-htmlx/internal/app/service"
)

func main() {

	// Initialize your services here
	contentService := service.NewContentService()

	// Initialize the handler with the content service
	contentHandler := handler.NewHandler(contentService)

	// Set up the HTTP server with the necessary routes
	router := web.NewRouter()
	router.HandleFunc("/", contentHandler.ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/update", contentHandler.ServeHTTP).Methods(http.MethodGet)

	router.HandleFunc("/readme", contentHandler.ServeHTTP).Methods(http.MethodGet)

	absPublicPath, _ := filepath.Abs("public")
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(absPublicPath))))

	absStaticPath, _ := filepath.Abs("static")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(absStaticPath))))

	// Start the server
	port := ":9090"
	log.Printf("Starting server on %s\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
