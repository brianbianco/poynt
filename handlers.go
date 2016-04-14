package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home Handler...")
	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]string, 2)
	data["service"] = "Poynt"
	data["version"] = "1"
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Write(j)
	return
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	fmt.Println("Handling post for", vars["key"])
	if r.Body == nil {
		fmt.Println("The body of the request is empty dude")
	}

	decoder := json.NewDecoder(r.Body)

	for {
		var jp JsonPoynt
		if err := decoder.Decode(&jp); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		jp.Name = vars["key"]
		// return the data for testing :)
		w.Write(jp.Data)
		poynts.Write(vars["key"], jp.ToPoynt())
	}
	return
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	if _, ok := r.URL.Query()["t_log"]; ok {
	}

	var data []JsonPoynt
	for _, v := range poynts.Filter(vars["key"], r.URL.Query()) {
		data = append(data, v.ToJson())
	}
	fmt.Println("Trying to marshal", data)
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(j)
	return
}
