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
	sk, sk1 := KeypairAgree(alice, bob.PublicKey.Bytes()), KeypairAgree(bob, alice.PublicKey.Bytes())

	if !bytes.Equal(sk, sk1) {
		t.Errorf("Shared keys don't match: sk=%d, sk1=%d", sk, sk1)
	}

	AliceInit(&aliceState, sk, bob.PublicKey.Bytes())

	if aliceState.DHs.Bytes() == nil {
		t.Errorf("aliceState.DHs is nil!")
	}
}
