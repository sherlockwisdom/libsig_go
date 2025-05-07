package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
	"encoding/binary"
)

type Headers struct {
	Dh Keypairs
	Pn uint32
	N uint32
}

type States struct {
	DHs Keypairs
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

func (s States) GobEncode() ([]byte, error) {
	var mkSkippedBuffer bytes.Buffer
	enc := gob.NewEncoder(&mkSkippedBuffer)
	err := enc.Encode(s.MKSKIPPED)
	if err != nil {
		panic(err)
	}

	// Assuming everything is 32 bytes
	Len := (32 * 6) + (4*3) + len(mkSkippedBuffer.Bytes())

	result := make([]byte, Len)
	copy(result[0:32], s.DHs.PrivateKey.Bytes())
	copy(result[32:64], s.DHs.PrivateKey.PublicKey.Bytes())

	copy(result[64:96], s.DHr)
	copy(result[96:128], s.RK)
	copy(result[128:160], s.CKs)
	copy(result[160:192], s.CKr)

	binary.BigEndian.PutUint32(result[192:196], s.Ns)
	binary.BigEndian.PutUint32(result[196:200], s.Nr)
	binary.BigEndian.PutUint32(result[200:204], s.PN)
	copy(result[204:(204 + len(mkSkippedBuffer.Bytes()))], mkSkippedBuffer.Bytes())

	if len(result) != (204 + len(mkSkippedBuffer.Bytes())) {
		return result, fmt.Errorf(
			"Serializaton lengths not equal: wanted: %d, got: %d",
			Len,
			(204 + len(mkSkippedBuffer.Bytes())),
		)
	}

	return result, nil 
}

func (s *States) GobDecode(data []byte) error {
	if len(data) < (32 * 6) + (4*3) {
		return fmt.Errorf(
			"Not enough data to build state: wanted: %d, got: %d",
			(32 * 6) + (4*3),
			len(data),
		)
	}

	DHsPrivateKey := make([]byte, 32)
	DHsPublicKey := make([]byte, 32)
	copy(DHsPrivateKey, data[0:32])
	copy(DHsPublicKey, data[32:64])

	s.DHs.PrivateKey.SetBytes(DHsPrivateKey)
	s.DHs.PrivateKey.PublicKey.SetBytes(DHsPublicKey)

	s.DHr = make([]byte, len(data[64:96]))
	copy(s.DHr, data[64:96])

	s.RK = make([]byte, len(data[96:128]))
	copy(s.RK, data[96:128])

	s.CKs = make([]byte, len(data[128:160]))
	copy(s.CKs, data[128:160])

	s.CKr = make([]byte, len(data[160:192]))
	copy(s.CKr, data[160:192])

	s.Ns = binary.BigEndian.Uint32(data[192:196])
	s.Nr = binary.BigEndian.Uint32(data[196:200])
	s.PN = binary.BigEndian.Uint32(data[200:204])

	dec := gob.NewDecoder(bytes.NewReader(data[204:]))
	err := dec.Decode(&s.MKSKIPPED)

	return err
}

func (m States) Serialize() (bytes.Buffer, error) {
	var buffer bytes.Buffer

	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(m)

	if err != nil {
		panic(err)
	}

	return buffer, err
}

func (m *States) Deserialize(buffer bytes.Buffer) error {
	dec := gob.NewDecoder(&buffer)
	err := dec.Decode(&m)

	if err != nil {
		panic(err)
	}

	return err
}

func (m Headers) Serialize() ([]byte, error) {
	result := make([]byte, (4*2) + len(m.Dh.PrivateKey.PublicKey.Bytes()))
	binary.LittleEndian.PutUint32(result[0:4], m.N)
	binary.LittleEndian.PutUint32(result[4:8], m.Pn)
	copy(result[8:], m.Dh.PrivateKey.PublicKey.Bytes())
	return result, nil
}

func (m *Headers) Deserialize(buffer []byte) error {
	if len(buffer) < 8 {
		return fmt.Errorf("buffer too short: need at least 8 bytes, got %d", len(buffer))
	}

	m.N = binary.LittleEndian.Uint32(buffer[0:4])
	m.Pn = binary.LittleEndian.Uint32(buffer[4:8])
	pubKey := make([]byte, len(buffer)-8)
	copy(pubKey, buffer[8:])
	m.Dh.PrivateKey.PublicKey.SetBytes(pubKey)

	return nil
}
