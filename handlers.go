package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	fmt.Println("Running get handler")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	keyspace := poynts.Get(vars["key"])

	//Right now no forced ordering of shaping functions
	//This should be changed
	//Filter should always occur before sort and limit
	for param, value := range r.URL.Query() {
		switch param {
		case "sort_by":
			parsed := SortParser(value[0])
			sorts := make([]lessFunc, 0, len(parsed))
			for i, v := range parsed {
				sorts = append(sorts, SortMapper(i, v))
			}
			keyspace.Sort(sorts)
		case "limit":
			lim, _ := strconv.Atoi(value[0])
			keyspace.Limit(lim)
			// Assume the param is a supported filter
		case "col":
			cols := strings.Split(value[0], ",")
			keyspace.Select(cols)
		default:
			s, _ := CheckAndPadDate(value[0])
			keyspace.Filter(param, s)
		}
	}

	var data []JsonPoynt
	// Get the results and convert to Json
	for _, v := range keyspace.Apply() {
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

func PoyntHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running Poynt Handler")

	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fmt.Println(vars)

	//	var data []JsonPoynt
	keyspace := poynts.Get(vars["key"])
	p := []string{"log", "obs", "op"}

	for _, k := range p {
		if _, ok := vars[k]; ok {
			keyspace.Filter(k, vars[k])
		}
	}

	var data []JsonPoynt
	// Get the results and convert to Json

	for _, v := range keyspace.Apply() {
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

func ListKeysHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Running ListKeys handler")
	w.Header().Set("Content-Type", "application/json")
	keys := poynts.List()
	j, err := json.Marshal(keys)
	if err != nil {
		fmt.Println(err)
	}
	w.Write(j)
	return
}
