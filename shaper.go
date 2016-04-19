package main

type shaper interface {
	Filter(name string, value string) shaper
	Sort([]lessFunc) shaper
	Limit(l int) shaper
	Apply() []Poynt
}
