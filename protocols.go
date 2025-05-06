package main

import (
	"bytes"
	"encoding/gob"
)

type States struct {
	DHs []byte 
	DHr []byte

	RK []byte
	CKs []byte
	CKr []byte

	Ns int
	Nr int

	PN int

	MKSKIPPED map[string]int
}

func SerializeStates(state States) (bytes.Buffer, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)

	state = States{
		state.DHs,
		state.DHr,
		state.RK,
		state.CKs,
		state.CKr,
		state.Ns,
		state.Nr,
		state.PN,
		state.MKSKIPPED,
	}
	err := enc.Encode(state)

	return buffer, err
}

func DeserializeStates(statesBuffer bytes.Buffer) (States, error) {
	dec := gob.NewDecoder(&statesBuffer)

	var states States
	err := dec.Decode(&states)

	return states, err
}
