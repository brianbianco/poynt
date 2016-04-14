package main

import (
	"encoding/json"
	"errors"
	"time"
)

type Poynt struct {
	Name string
	Logt time.Time
	Obst time.Time
	Opt  time.Time
	Data json.RawMessage
}

type compare func(t time.Time) bool

func (p *Poynt) ToJson() JsonPoynt {
	var j JsonPoynt
	j.Name = p.Name
	j.Logt = TimeToString(p.Logt)
	j.Obst = TimeToString(p.Obst)
	j.Opt = TimeToString(p.Opt)
	j.Data = p.Data
	return j
}

func (p *Poynt) Compare(o string, t time.Time) (bool, error) {
	f := map[string]compare{
		"log":    p.Log,
		"log_lt": p.Log_lt,
		"log_le": p.Log_le,
		"log_gt": p.Log_gt,
		"log_ge": p.Log_ge,
		"obs":    p.Obs,
		"obs_lt": p.Obs_lt,
		"obs_le": p.Obs_le,
		"obs_gt": p.Obs_gt,
		"obs_ge": p.Obs_ge,
		"op":     p.Op,
		"op_lt":  p.Op_lt,
		"op_le":  p.Op_le,
		"op_gt":  p.Op_gt,
		"op_ge":  p.Op_ge,
	}
	if c, ok := f[o]; ok {
		return c(t), nil
	} else {
		err := errors.New("Poynt: Invalid comparison operator")
		return false, err
	}
}

// Log comparison functions
func (p *Poynt) Log(t time.Time) bool {
	return p.Logt.Unix() == t.Unix()
}

func (p *Poynt) Log_lt(t time.Time) bool {
	return p.Logt.Unix() < t.Unix()
}

func (p *Poynt) Log_le(t time.Time) bool {
	return p.Logt.Unix() <= t.Unix()
}

func (p *Poynt) Log_gt(t time.Time) bool {
	return p.Logt.Unix() > t.Unix()
}

func (p *Poynt) Log_ge(t time.Time) bool {
	return p.Logt.Unix() >= t.Unix()
}

// Obst comparison functions
func (p *Poynt) Obs(t time.Time) bool {
	return p.Obst.Unix() == t.Unix()
}

func (p *Poynt) Obs_lt(t time.Time) bool {
	return p.Obst.Unix() < t.Unix()
}

func (p *Poynt) Obs_le(t time.Time) bool {
	return p.Obst.Unix() <= t.Unix()
}

func (p *Poynt) Obs_gt(t time.Time) bool {
	return p.Obst.Unix() > t.Unix()
}

func (p *Poynt) Obs_ge(t time.Time) bool {
	return p.Obst.Unix() >= t.Unix()
}

// Opt time related comparison functions
func (p *Poynt) Op(t time.Time) bool {
	return p.Opt.Unix() == t.Unix()
}

func (p *Poynt) Op_lt(t time.Time) bool {
	return p.Opt.Unix() < t.Unix()
}

func (p *Poynt) Op_le(t time.Time) bool {
	return p.Opt.Unix() <= t.Unix()
}

func (p *Poynt) Op_gt(t time.Time) bool {
	return p.Opt.Unix() > t.Unix()
}

func (p *Poynt) Op_ge(t time.Time) bool {
	return p.Opt.Unix() >= t.Unix()
}
