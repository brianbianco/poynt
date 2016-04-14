package main

import (
	"encoding/json"
	"time"
)

type JsonPoynt struct {
	Name string
	Logt string          `json:"t_log"`
	Obst string          `json:"t_obs"`
	Opt  string          `json:"t_op"`
	Data json.RawMessage `json:"data"`
}

type Poynt struct {
	Name string
	Logt time.Time
	Obst time.Time
	Opt  time.Time
	Data json.RawMessage
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

func (p *Poynt) ToJson() JsonPoynt {
	var j JsonPoynt
	j.Name = p.Name
	j.Logt = TimeToString(p.Logt)
	j.Obst = TimeToString(p.Obst)
	j.Opt = TimeToString(p.Opt)
	j.Data = p.Data
	return j
}
