package main

import (
	"maze.io/x/crypto/x25519"
	"golang.org/x/crypto/curve25519"
	"crypto/rand"
)

type KeypairsInterface interface {
	Init() 
	GetPublicKey([]byte) 
	Agree([]byte) []byte
}

type Keypairs struct {
	PrivateKey x25519.PrivateKey
}

func (k *Keypairs) Init() {
	privateKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	k.PrivateKey = *privateKey
}

func (k *Keypairs) GetPublicKey() {
	publicKey, err := curve25519.X25519(k.PrivateKey.Bytes(), curve25519.Basepoint)
	if err != nil {
		panic(err)
	}
	k.PrivateKey.PublicKey.SetBytes(publicKey)
}


func (k Keypairs) Agree(peerPublicKeyRaw []byte) []byte {
	var peerPublicKey x25519.PublicKey 
	peerPublicKey.SetBytes(peerPublicKeyRaw)

	return k.PrivateKey.Shared(&peerPublicKey)
}
