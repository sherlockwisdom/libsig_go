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

	var alice, bob Keypairs
	alice.Init()
	bob.Init()

	sk := alice.Agree(bob.PrivateKey.PublicKey.Bytes())
	sk1 := bob.Agree(alice.PrivateKey.PublicKey.Bytes())

	if !bytes.Equal(sk, sk1) {
		t.Errorf("Shared keys don't match: sk=%d, sk1=%d", sk, sk1)
	}

	AliceInit(&aliceState, sk, bob.PrivateKey.PublicKey.Bytes())

	if aliceState.DHs.PrivateKey.Bytes() == nil {
		t.Errorf("aliceState.DHs is nil!")
	}
}
