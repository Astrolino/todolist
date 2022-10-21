package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Healthz called")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}
func main() {
	log.Println("Starting server")
	router := mux.NewRouter()
	router.HandleFunc("/healthz", Healthz).Methods("GET")
	http.ListenAndServe(":8000", router)
}
