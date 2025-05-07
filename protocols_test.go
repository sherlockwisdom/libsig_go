package main

import (
	"testing"
	"reflect"
	"bytes"
	"math/rand"
	"encoding/base64"
)

func TestSK(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init()

	var keypairs1 Keypairs
	keypairs1.Init()

	if bytes.Equal(
		keypairs.PrivateKey.PublicKey.Bytes(), 
		keypairs1.PrivateKey.PublicKey.Bytes(),
	) {
		t.Errorf(
			"PK are equal: got %d, %d", 
			keypairs.PrivateKey.Bytes(), 
			keypairs1.PrivateKey.Bytes(),
		)
	}

	sk := keypairs.Agree(keypairs1.PrivateKey.PublicKey.Bytes())
	sk1 := keypairs1.Agree(keypairs.PrivateKey.PublicKey.Bytes())

	if bytes.Equal(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
}

func TestHeadersEncodeDecode(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init()

	var header = Headers {
		keypairs,
		0,
		0,
	}

	serializedHeader, err := header.Serialize()
	if err != nil {
		t.Errorf("Error serializing header: %s", err)
	}

	var header1 Headers
	err = header1.Deserialize(serializedHeader)
	if err != nil {
		t.Errorf("Error deserializing header: %s", err)
	}

	if !bytes.Equal(
		header1.Dh.PrivateKey.PublicKey.Bytes(), 
		header.Dh.PrivateKey.PublicKey.Bytes(), 
	) {
		t.Errorf("Headers do not match: wanted: %d, got: %d", header.Dh, header1.Dh)
	}
}

func TestStatesEncodeDecode(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init()

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
		keypairs,
		DHr,
		RK,
		CKs,
		CKr,
		0,
		0,
		0,
		MKSKIPPED,
	}

	serializedState, err := state.Serialize()
	if err != nil {
		t.Errorf("Error serializing: %s", err)
	}

	var state1 States
	err = state1.Deserialize(serializedState)
	if err != nil {
		t.Errorf("Error deserializing: %s", err)
	}

	if reflect.DeepEqual(state1, state) == false {
		if !bytes.Equal(state.DHs.PrivateKey.Bytes(), state1.DHs.PrivateKey.Bytes()) {
			t.Errorf(
				"States do not match: DHs.Private key wanted: %d, got: %d",
				state.DHs.PrivateKey.Bytes(),
				state1.DHs.PrivateKey.Bytes(),
			)
		}
		if !bytes.Equal(state.DHs.PrivateKey.PublicKey.Bytes(), state1.DHs.PrivateKey.PublicKey.Bytes()) {
			t.Errorf(
				"States do not match: DHs.Public key wanted: %d, got: %d",
				state.DHs.PrivateKey.PublicKey.Bytes(),
				state1.DHs.PrivateKey.PublicKey.Bytes(),
			)
		}
		if !bytes.Equal(state.DHr, state1.DHr) {
			t.Errorf("States do not match: DHr key wanted: %d, got: %d", state.DHr, state1.DHr)
		}
		t.Errorf("States do not match...")
	}

	/*
	deserializedState.Ns = state.Ns + 1

	if reflect.DeepEqual(deserializedState, state) {
		t.Errorf("States match when should not: wanted Ns:=%d, got Ns:=%d", state.Ns, deserializedState.Ns)
	}
	*/
}
