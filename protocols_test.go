package main

import (
	"testing"
	"reflect"
	"bytes"
	"math/rand"
	"encoding/base64"
)

func TestSK(t *testing.T) {
	pk := KeypairInit()
	pk1 := KeypairInit()

	if bytes.Equal(pk.PublicKey.Bytes(), pk1.PublicKey.Bytes()) {
		t.Errorf("PK are equal: got %d, %d", pk.Bytes(), pk1.Bytes())
	}

	sk := KeypairAgree(pk.Bytes(), pk1.PublicKey.Bytes())
	sk1 := KeypairAgree(pk1.Bytes(), pk.PublicKey.Bytes())

	if bytes.Equal(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
}

func TestHeadersEncodeDecode(t *testing.T) {
	DH := make([]byte, 32)
	rand.Read(DH)

	header := Headers {
		DH,
		0,
		0,
	}

	serializedHeader, err := SerializeProtocols(header)
	if err != nil {
		t.Errorf("Error serializing header: %s", err)
	}

	deserializedHeader, err := DeserializeProtocols[Headers](serializedHeader)
	if err != nil {
		t.Errorf("Error deserializing header: %s", err)
	}

	if reflect.DeepEqual(deserializedHeader, header) == false {
		t.Errorf("Headers do not match...")
	}
}

func TestStatesEncodeDecode(t *testing.T) {
	DHs := make([]byte, 32)
	rand.Read(DHs)

	DHr := make([]byte, 32)
	rand.Read(DHr)

	RK := make([]byte, 32)
	rand.Read(RK)

	CKs := make([]byte, 32)
	rand.Read(CKs)

	CKr := make([]byte, 32)
	rand.Read(CKr)

	pK := make([]byte, 32)
	rand.Read(pK)
	PK := base64.StdEncoding.EncodeToString(pK)

	var MKSKIPPED map[string]int
	MKSKIPPED = make(map[string]int)
	MKSKIPPED[PK] = 0

	state := States{
		DHs,
		DHr,
		RK,
		CKs,
		CKr,
		0,
		0,
		0,
		MKSKIPPED,
	}

	serializedState, err := SerializeProtocols(state)
	if err != nil {
		t.Errorf("Error serializing: %s", err)
	}

	deserializedState, err := DeserializeProtocols[States](serializedState)
	if err != nil {
		t.Errorf("Error deserializing: %s", err)
	}

	if reflect.DeepEqual(deserializedState, state) == false {
		t.Errorf("States do not match...")
	}

	deserializedState.Ns = state.Ns + 1

	if reflect.DeepEqual(deserializedState, state) {
		t.Errorf("States match when should not: wanted Ns:=%d, got Ns:=%d", state.Ns, deserializedState.Ns)
	}

	// fmt.Printf("%d\n", serializedState)
}
