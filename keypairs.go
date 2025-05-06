package main

import (
	"maze.io/x/crypto/x25519"
	"golang.org/x/crypto/curve25519"
	"crypto/rand"
)

func KeypairInit() x25519.PrivateKey {
	privateKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	return *privateKey
}

func KeypairGetPublicKey(privateKey []byte) []byte {
	publicKey, err := curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		panic(err)
	}
	return publicKey
}


func KeypairAgree(privateKeyRaw, peerPublicKeyRaw []byte) []byte {
	var peerPublicKey x25519.PublicKey 
	peerPublicKey.SetBytes(peerPublicKeyRaw)

	var privateKey x25519.PrivateKey
	privateKey.SetBytes(privateKeyRaw)
	skSlice := privateKey.Shared(&peerPublicKey)

	// return *(*[32]byte)(skSlice) // return in arrays not slices
	return skSlice
}
