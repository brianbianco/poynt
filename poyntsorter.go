package main

import (
	"sort"
)

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

// Our sorting functions, put here for convenience
func AscLogt(p1 *Poynt, p2 *Poynt) bool {
	return p1.Log_lt(p2.Logt)
}
func AscObst(p1 *Poynt, p2 *Poynt) bool {
	return p1.Obs_lt(p2.Obst)
}
func AscOpt(p1 *Poynt, p2 *Poynt) bool {
	return p1.Op_lt(p2.Opt)
}

func DscLogt(p1 *Poynt, p2 *Poynt) bool {
	return p1.Log_gt(p2.Logt)
}
func DscObst(p1 *Poynt, p2 *Poynt) bool {
	return p1.Obs_gt(p2.Obst)
}
func DscOpt(p1 *Poynt, p2 *Poynt) bool {
	return p1.Op_gt(p2.Opt)
}
