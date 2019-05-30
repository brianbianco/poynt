package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var poynts PoyntStore

func main() {
	fmt.Println("Poynt service")
	poynts = NewMemoryPoyntStore()
	//poynts = NewS3PoyntStore("cargometrics-infra", "poynts_test", "us-west-2")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{apiVersion}/poynt/keys", ListKeysHandler)
	r.HandleFunc("/{apiVersion}/poynt/{key}", PostHandler).Methods("POST")
	r.HandleFunc("/{apiVersion}/poynt/{key}", GetHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}", PoyntHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}/{obs}", PoyntHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}/{obs}/{op}", PoyntHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
