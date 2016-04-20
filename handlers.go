package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//Maps x,x,x where O is the order to Logical, Observational, Operational
//where x is a or d for Ascending and Descending order
var SortMap = [3]map[string]lessFunc{
	map[string]lessFunc{"a": AscLogt, "d": DscLogt},
	map[string]lessFunc{"a": AscObst, "d": DscObst},
	map[string]lessFunc{"a": AscOpt, "d": DscOpt},
}

func SortMapper(position int, ordering string) lessFunc {
	return SortMap[position][ordering]
}

func SortParser(input string) []string {
	max := 3
	results := make([]string, 0, max)
	myExp := regexp.MustCompile(`([d|a])(\d*)(?:,{0,1})`)
	match := myExp.FindAllStringSubmatch(input, -1)
	for i := 0; i < max; i++ {
		//first field is the fully matched string, ignore it.
		//second field is a or d
		results = append(results, match[i][1])
		//third field is a number, currently not used
	}
	return results
}

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

	var data []JsonPoynt
	keyspace := poynts.Get(vars["key"])

	// Call a filter for each parameter
	// This will have to be changed to handle more than just
	// the list of comparison methods implemented by the poynt struct
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
		default:
			keyspace.Filter(param, value[0])
		}
	}

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
