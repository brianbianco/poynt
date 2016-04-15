package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

type MemoryPoyntStore struct {
	store map[string]map[string]Poynt
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

func (s *MemoryPoyntStore) Filter(key string, f map[string][]string) []Poynt {
	var pa []Poynt
	if _, ok := s.store[key]; !ok {
		return []Poynt{}
	}
	for _, p := range s.store[key] {
		matches := true
		for param, value := range f {
			t, _ := StringToTime(value[0])
			matches, _ = p.Compare(param, t)
			if !matches {
				break
			}
		}
		if matches {
			pa = append(pa, p)
		}
	}

	// Our two sorting functions we want to order by
	logt := func(p1 *Poynt, p2 *Poynt) bool {
		return p1.Log_lt(p2.Logt)
	}
	obst := func(p1 *Poynt, p2 *Poynt) bool {
		return p1.Obs_lt(p2.Obst)
	}

	OrderedBy(logt, obst).Sort(pa)
	return pa
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

/* Lets make poynts sortable

logt := func(p1 *Poynt, p2 *Poynt) bool {
        return p1.Log_lt(p2.Log)
}
*/

//type By func(p1 *Poynt, p2 *Poynt) bool
type lessFunc func(p1 *Poynt, p2 *Poynt) bool

func OrderedBy(less ...lessFunc) *poyntSorter {
	return &poyntSorter{
		less: less,
	}
}

func (ps *poyntSorter) Sort(poynts []Poynt) {
	ps.poynts = poynts
	sort.Sort(ps)
}

type poyntSorter struct {
	poynts []Poynt
	less   []lessFunc
}

func (ps *poyntSorter) Len() int {
	return len(ps.poynts)
}

func (ps *poyntSorter) Swap(i int, j int) {
	ps.poynts[i], ps.poynts[j] = ps.poynts[j], ps.poynts[i]
}

func (ps *poyntSorter) Less(i int, j int) bool {
	p1, p2 := &ps.poynts[i], &ps.poynts[j]
	var k int
	for k = 0; k < len(ps.less)-1; k++ {
		less := ps.less[k]
		switch {
		case less(p1, p2):
			return true
		case less(p2, p1):
			return false
		}
	}
	// all other comparisons claim they are equal so lets go ahead and
	// return the results of the final one
	return ps.less[k](p1, p2)
}
