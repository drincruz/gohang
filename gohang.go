package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func getDefaultPort() string {
	port, exists := os.LookupEnv("GOHANG_PORT")
	if exists {
		return fmt.Sprintf(":%s", port)
	}
	return ":5000"
}

func writeJsonResponseHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func OkTwoHundred(w http.ResponseWriter, r *http.Request) {
	writeJsonResponseHeader(w)
	w.Write([]byte("{ \"data\": \"200 OK\" }"))
}

func FiveHundred(w http.ResponseWriter, r *http.Request) {
	writeJsonResponseHeader(w)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{ \"data\": \"500 Internal Server Error\" }"))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	writeJsonResponseHeader(w)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("{ \"data\": \"404 Not Found\" }"))
}

func SlowResponse(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	writeJsonResponseHeader(w)
	w.Write([]byte("{ \"data\": \"slow response\"}"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", OkTwoHundred)
	mux.HandleFunc("/500", FiveHundred)
	mux.HandleFunc("/404", NotFound)
	mux.HandleFunc("/slow", SlowResponse)

	port := getDefaultPort()

	log.Printf("[INFO] Now listening on %s", port)
	http.ListenAndServe(port, mux)
}
