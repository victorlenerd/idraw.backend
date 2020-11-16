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

	router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, req)
		})
	})

	//router.HandleFunc("/", handlers.IndexHandler).Methods(http.MethodGet)
	router.HandleFunc("/upload/{noteID}", handlers.UploadHandler).Methods(http.MethodPost)
	router.HandleFunc("/notes/{noteID}", handlers.GetNoteImages).Methods(http.MethodGet)
	router.HandleFunc("/notes/{noteID}", handlers.CreateNote).Methods(http.MethodPost)

	fileServer := http.FileServer(http.Dir("./public/"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer)).Methods(http.MethodGet)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}