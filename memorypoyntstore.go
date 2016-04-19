package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type MemoryPoyntStore struct {
	store map[string]map[string]Poynt
}

func (s *MemoryPoyntStore) Get(key string) shaper {
	var pc PoyntCollection
	if _, ok := s.store[key]; !ok {
		return &pc
	} else {
		pc.store = mapToArray(s.store[key])
		return &pc
	}
}

func (s *MemoryPoyntStore) Write(key string, p Poynt) bool {
	id := uniqueId(p)
	if namespace, ok := s.store[key]; ok {
		if _, ok := namespace[id]; ok {
			fmt.Println("Error this poynt already exists")
			return false
		} else {
			namespace[id] = p
		}
	} else {
		s.store[key] = map[string]Poynt{id: p}
	}
	return true
}

func mapToArray(m map[string]Poynt) []Poynt {
	results := make([]Poynt, len(m))
	var i = 0
	for _, p := range m {
		results[i] = p
		i++
	}
	return results
}

func uniqueId(p Poynt) string {
	jp := p.ToJson()
	h := sha1.New()
	h.Write([]byte(jp.Opt + jp.Logt + jp.Obst))
	return hex.EncodeToString(h.Sum(nil))
}
