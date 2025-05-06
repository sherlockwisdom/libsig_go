package main

import (
	"maze.io/x/crypto/x25519"
	"crypto/rand"
)

func Init() (x25519.PublicKey, []byte, error) {
	pk, err := x25519.GenerateKey(rand.Reader)

	// TODO: store the private key

	return pk.PublicKey, pk.Bytes(), err 
}

// func Agree(peerPublicKeyRaw, privateKeyRaw []byte) [32]byte { // If needed in arrays not slices
func Agree(peerPublicKeyRaw, privateKeyRaw []byte) []byte {
	var peerPublicKey x25519.PublicKey 
	var privateKey x25519.PrivateKey 

	peerPublicKey.SetBytes(peerPublicKeyRaw)
	privateKey.SetBytes(privateKeyRaw)

	// return privateKey.Shared(&peerPublicKey)
	skSlice := privateKey.Shared(&peerPublicKey)

	// return *(*[32]byte)(skSlice) // return in arrays not slices
	return skSlice
}
