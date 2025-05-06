package main

import (
	"testing"
	"bytes"
	/*
	"fmt"
	"reflect"
	"math/rand"
	"encoding/base64"
	*/
)

func TestRatchets(t *testing.T) {
	var aliceState States

	alice, bob := KeypairInit(), KeypairInit()
	sk := KeypairAgree(alice.Bytes(), bob.PublicKey.Bytes())
	sk1 := KeypairAgree(bob.Bytes(), alice.PublicKey.Bytes())

	if !bytes.Equal(sk, sk1) {
		t.Errorf("Shared keys don't match: sk=%d, sk1=%d", sk, sk1)
	}

	AliceInit(&aliceState, sk, bob.PublicKey.Bytes())

	if aliceState.DHs == nil {
		t.Errorf("aliceState.DHs is nil!")
	}
}
