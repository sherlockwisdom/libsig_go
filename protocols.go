package main

import (
	"bytes"
	"encoding/gob"
	"maze.io/x/crypto/x25519"
)

type Headers struct {
	Dh []byte
	Pn int
	N int
}

type States struct {
	DHs x25519.PrivateKey
	DHr []byte

	RK []byte
	CKs []byte
	CKr []byte

	Ns int
	Nr int

	PN int

	MKSKIPPED map[string]int
}

type Protocol interface {
	Headers | States
}

func SerializeProtocols[V Protocol](protocol V) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(protocol)

	return buffer, err
}

func DeserializeProtocols[V Protocol](buffer bytes.Buffer) (V, error) {
	dec := gob.NewDecoder(&buffer)

	var protocol V
	err := dec.Decode(&protocol)

	return protocol, err
}
