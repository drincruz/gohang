package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
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

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	writeJsonResponseHeader(w)
	request, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatalf("[EchoHandler] Error happened with DumpRequest. Err: %s", err)
	}
	responseMap := make(map[string]string)
	responseMap["data"] = string(request)
	response, err := json.Marshal(responseMap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("[EchoHandler] Error happened with JSON marshal. Err: %s", err)
		w.Write([]byte(fmt.Sprintf("{\"status\": \"error\", \"response\": \"%s\"}", err)))
		return
	}
	w.Write([]byte(response))
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
	mux.HandleFunc("/echo", EchoHandler)
	mux.HandleFunc("/slow", SlowResponse)

	port := getDefaultPort()

	log.Printf("[INFO] Now listening on %s", port)
	http.ListenAndServe(port, mux)
}
