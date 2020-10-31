package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"idraw/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.Use(mux.CORSMethodMiddleware(router))

	router.HandleFunc("/", handlers.IndexHandler).Methods(http.MethodGet)
	router.HandleFunc("/upload/{noteID}", handlers.UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/notes/{noteID}", handlers.GetNoteImages).Methods(http.MethodGet)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}