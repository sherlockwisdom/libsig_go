package main

import (
	"bytes"
	"testing"
	/*
		"fmt"
		"reflect"
		"math/rand"
		"encoding/base64"
	*/)

func TestRatchets(t *testing.T) {
	var aliceState States

	var alice, bob Keypairs
	alice.Init("db/alice.db")
	bob.Init("db/bob.db")

	sk, err := alice.Agree(bob.PrivateKey.PublicKey.Bytes())
	if err != nil {
		t.Error(err)
	}

	sk1, err := bob.Agree(alice.PrivateKey.PublicKey.Bytes())
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(sk, sk1) {
		t.Errorf("Shared keys don't match: sk=%d, sk1=%d", sk, sk1)
	}

	AliceInit(&aliceState, sk, bob.PrivateKey.PublicKey.Bytes(), "db/alice_ratchet.db")

	if aliceState.DHs.PrivateKey.Bytes() == nil {
		t.Errorf("aliceState.DHs is nil!")
	}
}
