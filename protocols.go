package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"encoding/binary"
)

type Headers struct {
	Dh []byte
	Pn uint32
	N uint32
}

type States struct {
	DHs []byte
	DHr []byte

	RK []byte
	CKs []byte
	CKr []byte

	Ns uint32
	Nr uint32

	PN uint32

	MKSKIPPED map[string]int
}

type ProtocolStates interface {
	Serialize() (bytes.Buffer, error)
	Deserialize(bytes.Buffer) error
}

type ProtocolsHeaders interface {
	Serialize() ([]byte, error)
	Deserialize([]byte) error
}

func (m States) Serialize() (bytes.Buffer, error) {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(m)

	return buffer, err
}

func (m *States) Deserialize(buffer bytes.Buffer) error {
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&m)
	return err
}

func (m Headers) Serialize() ([]byte, error) {
	result := make([]byte, (4*2) + len(m.Dh))
	binary.LittleEndian.PutUint32(result[0:4], m.N)
	binary.LittleEndian.PutUint32(result[4:8], m.Pn)
	copy(result[8:], m.Dh)
	return result, nil
}

func (m *Headers) Deserialize(buffer []byte) error {
	if len(buffer) < 8 {
		return fmt.Errorf("buffer too short: need at least 8 bytes, got %d", len(buffer))
	}

	m.N = binary.LittleEndian.Uint32(buffer[0:4])
	m.Pn = binary.LittleEndian.Uint32(buffer[4:8])
	m.Dh = make([]byte, len(buffer)-8)
	copy(m.Dh, buffer[8:])

	return nil
}
