package main

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
	/*
		funcs := []lessFunc{DscObst, AscOpt}
		OrderedBy(funcs...).Sort(pa)
	*/
	return pc
}

func (pc *PoyntCollection) Limit(l int) shaper {
	return pc
}

func (pc *PoyntCollection) Apply() []Poynt {
	for _, f := range pc.shapers {
		f()
	}
	return pc.store
}
