package main

type PoyntStore interface {
	Write(key string, p Poynt) bool
	Get(key string) shaper
	List() []string
}
