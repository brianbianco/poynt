package main

type PoyntStore interface {
	Write(key string, p Poynt) bool

	/* Filter needs to implement the following predicates, they will be passed in
	   as a URL query string
	    op, op_lt, op_le, op_gt, op_ge
	    obs, obs_lt, obs_le, obs_gt, obs_ge
	    log, log_lt, log_le, log_gt, log_ge
	*/
	Filter(key string, f map[string][]string) []Poynt
}
