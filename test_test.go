package main


import (
	"testing"
	"bytes"
	/*
	"crypto/rand"
	"maze.io/x/crypto/x25519"
	"golang.org/x/crypto/curve25519"
	*/
)

func TestDeterministicPublicKeys(t *testing.T) {
	privKey := KeypairInit()
	value := privKey.Bytes()

	pubKey := KeypairGetPublicKey(value) 

	if !bytes.Equal(pubKey, privKey.PublicKey.Bytes()) {
		t.Errorf("Pub keys differ! \n\twanted %d, \n\tgot %d", value, pubKey)
	}
}
