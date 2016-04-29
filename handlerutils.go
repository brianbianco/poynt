package main

import (
	"errors"
	"regexp"
	"strings"
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

// Validates that passed in string conforms to date/time format
// YYYYmmddHHMMSS.mmm
// Will pad out with Zeros anything after YYYYmmdd
func CheckAndPadDate(s string) (string, error) {
	size := len(s)
	min := 8 // YYYYmmddd
	var err error

	switch {
	case size < min:
		err = errors.New("Parameter doens't meet min size requirement")
	case size > TimeFormatMaxLength:
		err = errors.New("Parameter exceeds max size")
	case size >= min && size < TimeFormatMaxLength:
		return PadOut(s), err
	}
	return s, err
}

// takes a string of format 20160320183001.001
// Pads out anything over minimum with 0's
func PadOut(s string) string {
	s = strings.Replace(s, ".", "", -1)
	size := len(s)
	// Reduce size by one because period was stripped out
	padding := TimeFormatMaxLength - size - 1
	if padding > 0 {
		fixed := s + strings.Repeat("0", padding)
		size = len(fixed)
		fixed = fixed[0:size-3] + "." + fixed[size-3:size]
		return fixed
	} else {
		return s
	}
}
