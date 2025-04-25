package main

import (
	"maze.io/x/crypto/x25519"
	"crypto/rand"
)

func Init() (x25519.PublicKey, error) {
	pk, err := x25519.GenerateKey(rand.Reader)

	// TODO: store the private key

	return pk.PublicKey, err 
}

func Agree(peerPublicKeyRaw, privateKeyRaw []byte) []byte {
	peerPublicKey = publicKey.SetBytes(peerPublicKeyRaw)
	privateKey = privateKey.SetBytes(privateKeyRaw)

	return privateKey.Shared(peerPublicKey)
}
