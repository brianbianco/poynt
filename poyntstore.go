package main

type PoyntStore interface {
	Write(key string, p Poynt) bool
	Filter(key string, f map[string][]string) []Poynt
}
