package main

import (
	"fmt"
	"testing"
	"reflect"
	"bytes"
	"math/rand"
	"encoding/base64"
)

func TestSK(t *testing.T) {
	pk, prKBytes, _ := Init()
	pk1, prK1Bytes, _ := Init()

	fmt.Println(reflect.TypeOf(pk.Bytes()))

	if pk == pk1 {
		t.Errorf("PK are equal: got %d, %d", pk.Bytes(), pk1.Bytes())
	}

	sk := Agree(pk1.Bytes(), prKBytes)
	sk1 := Agree(pk.Bytes(), prK1Bytes)

	/*
	if reflect.DeepEqual(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
	*/

	if bytes.Equal(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
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

	serializedState, err := SerializeStates(state)
	if err != nil {
		t.Errorf("Error serializing: %s", err)
	}

	deserializedState, err := DeserializeStates(serializedState)
	if err != nil {
		t.Errorf("Error deserializing: %s", err)
	}

	if reflect.DeepEqual(deserializedState, state) == false {
		t.Errorf("States do not match...")
	}

	fmt.Printf("%d\n", serializedState)
}
