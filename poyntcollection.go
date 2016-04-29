package main

import (
	"encoding/json"
)

type shaperFunc func() shaper

type PoyntCollection struct {
	store   []Poynt
	shapers []shaperFunc
}

func (pc *PoyntCollection) Filter(name string, value string) shaper {
	shaper := func() shaper {
		var pa []Poynt
		for _, p := range pc.store {
			t, _ := StringToTime(value)
			matches, _ := p.Compare(name, t)
			if matches {
				pa = append(pa, p)
			}
		}
		pc.store = pa
		return pc
	}
	pc.shapers = append(pc.shapers, shaper)
	return pc
}

func (pc *PoyntCollection) Sort(f []lessFunc) shaper {
	shaper := func() shaper {
		OrderedBy(f...).Sort(pc.store)
		return pc
	}
	pc.shapers = append(pc.shapers, shaper)
	return pc
}

func (pc *PoyntCollection) Select(cols []string) shaper {
	shaper := func() shaper {
		for k, p := range pc.store {
			selection := p.SelectCols(cols)
			reduced, _ := json.Marshal(selection)
			p.Data = (json.RawMessage)(reduced)
			pc.store[k] = p
		}
		return pc
	}
	pc.shapers = append(pc.shapers, shaper)
	return pc
}
func (pc *PoyntCollection) Limit(l int) shaper {
	shaper := func() shaper {
		if l > len(pc.store) {
			l = len(pc.store)
		}
		tmp := make([]Poynt, l, l)
		copy(tmp, pc.store)
		pc.store = tmp
		return pc
	}
	pc.shapers = append(pc.shapers, shaper)
	return pc
}

func (pc *PoyntCollection) Apply() []Poynt {
	for _, f := range pc.shapers {
		f()
	}
	return pc.store
}
