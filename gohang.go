package main

import (
	"log"
	"net/http"
	"time"
)

const DEFAULT_PORT string = ":5000"

func okTwoHundred(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{ \"data\": \"200 OK\"}"))
}

func fiveHundred(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{ \"data\": \"500 Internal Server Error\"}"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{ \"data\": \"404 Not Found\"}"))
}

func slowResponse(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.Write([]byte("{ \"data\": \"slow response\"}"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", okTwoHundred)
	mux.HandleFunc("/500", fiveHundred)
	mux.HandleFunc("/404", notFound)
	mux.HandleFunc("/slow", slowResponse)

	log.Printf("[INFO] Now listening on %s", DEFAULT_PORT)
	http.ListenAndServe(DEFAULT_PORT, mux)
}
