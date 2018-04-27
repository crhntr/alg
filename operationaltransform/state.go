package main

import "github.com/globalsign/mgo/bson"

type State struct {
	Text string `json:"text"`
}

type Operation struct {
	UserID bson.ObjectId `json:"uid"`
	StrID  int           `json:"sid"`
	Index  int           `json:"i"`
	Length string        `json:"l"`
	Value  string        `json:"v"`
	Type string           `json:"t"`
}

func InclusionTransformation(s State, a, b Operation) State {
	if a.Index < b.Index || a.Index == b.Index && a.ID < b.ID {
		return s.Text[:a.Index] + a.Value + s.Text[a.Index+a.Length:]
	}
	return s.Text[:a.Index+len(b.Value)] + a.Value + s.Text[a.Index+a.Length+len(b.Value):]
}

func ExclusionTransformation(s State, a, b Operation) State {
	if a.Index < b.Index || a.Index == b.Index && a.ID < b.ID {
		return s.Text[:a.Index] + a.Value + s.Text[a.Index+a.Length:]
	}
	return s.Text[:a.Index-len(b.Value)] + a.Value + s.Text[a.Index-a.Length+len(b.Value):]
}

func (state State) Key() string {
	return hex.EncodeToString(sha256.Sum256([]bytes(state.Text))))
}
