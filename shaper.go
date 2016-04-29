package main

// Shaper does NOT impose an order for its methods
type shaper interface {
	Select([]string) shaper
	Filter(name string, value string) shaper
	Sort([]lessFunc) shaper
	Limit(l int) shaper
	Apply() []Poynt
}
