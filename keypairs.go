package main

import (
	"maze.io/x/crypto/x25519"
	"crypto/rand"
)

func GetPublicKey() (*x25519.PrivateKey, error) {
	pk, err := x25519.GenerateKey(rand.Reader)
	return pk, err 
}
