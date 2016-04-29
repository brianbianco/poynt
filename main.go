package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var poynts MemoryPoyntStore

func main() {
	fmt.Println("Poynt service")
	poynts.store = make(map[string]map[string]Poynt)
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{apiVersion}/poynt/{key}", PostHandler).Methods("POST")
	r.HandleFunc("/{apiVersion}/poynt/{key}", GetHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}", PoyntHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}/{obs}", PoyntHandler).Methods("GET")
	r.HandleFunc("/{apiVersion}/poynt/{key}/{log}/{obs}/{op}", PoyntHandler).Methods("GET")

	http.ListenAndServe(":8080", r)
}
