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

	http.ListenAndServe(":8080", r)
}
