package main

import (
	"encoding/json"
)

type JsonPoynt struct {
	Name string
	Logt string          `json:"t_log"`
	Obst string          `json:"t_obs"`
	Opt  string          `json:"t_op"`
	Data json.RawMessage `json:"data"`
}

func (j *JsonPoynt) ToPoynt() Poynt {
	var p Poynt
	p.Name = j.Name
	p.Logt, _ = StringToTime(j.Logt)
	p.Obst, _ = StringToTime(j.Obst)
	p.Opt, _ = StringToTime(j.Opt)
	p.Data = j.Data
	return p
}
