package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"reflect"
	"testing"
)

func TestSK(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init("db/test_sk_alice.db")

	var keypairs1 Keypairs
	keypairs1.Init("db/test_sk_bob.db")

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

	sk, err := keypairs.Agree(keypairs1.PrivateKey.PublicKey.Bytes())
	if err != nil {
		t.Error(err)
	}

	sk1, err := keypairs1.Agree(keypairs.PrivateKey.PublicKey.Bytes())
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(sk, sk1) {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
}

func TestHeadersEncodeDecode(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init("db/test_header_bob.db")

	var header = Headers{
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
		t.Errorf(
			"Headers do not match: wanted: %d, got: %d",
			header.Dh.PrivateKey.PublicKey.Bytes(),
			header1.Dh.PrivateKey.PublicKey.Bytes(),
		)
	}
}

func TestStatesEncodeDecode(t *testing.T) {
	var keypairs Keypairs
	keypairs.Init("db/test_states.db")

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

	var MKSKIPPED = make(map[string]int)
	MKSKIPPED[PK] = 10

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

	if !reflect.DeepEqual(state1, state) {
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
		if !bytes.Equal(state.RK, state1.RK) {
			t.Errorf("States do not match: RK key wanted: %d, got: %d", state.RK, state1.RK)
		}
		if !bytes.Equal(state.CKs, state1.CKs) {
			t.Errorf("States do not match: CKs key wanted: %d, got: %d", state.CKs, state1.CKs)
		}
		if !bytes.Equal(state.CKr, state1.CKr) {
			t.Errorf("States do not match: CKr key wanted: %d, got: %d", state.CKr, state1.CKr)
		}
		if state.Ns != state1.Ns {
			t.Errorf("States do not match: Ns wanted: %d, got: %d", state.Ns, state1.Ns)
		}
		if state.Nr != state1.Nr {
			t.Errorf("States do not match: Nr wanted: %d, got: %d", state.Nr, state1.Nr)
		}
		if state.PN != state1.PN {
			t.Errorf("States do not match: PN wanted: %d, got: %d", state.PN, state1.PN)
		}
		if !reflect.DeepEqual(state.MKSKIPPED, state1.MKSKIPPED) {
			t.Errorf(
				"States do not match: MKSKIPPED wanted: %d, got: %d",
				state.MKSKIPPED[PK],
				state1.MKSKIPPED[PK],
			)
		}
		if state.DHs.KeystorePath != state1.DHs.KeystorePath {
			t.Errorf(
				"States do not match: KeystorePath wanted: %s, got: %s",
				state.DHs.KeystorePath,
				state1.DHs.KeystorePath,
			)
		}
		if state.DHs.Pnt != state1.DHs.Pnt {
			t.Errorf(
				"States do not match: Pnt wanted: %s, got: %s",
				state.DHs.Pnt,
				state1.DHs.Pnt,
			)
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
