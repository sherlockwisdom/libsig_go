package main

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/curve25519"
	"maze.io/x/crypto/x25519"
)

type KeypairsInterface interface {
	Init()
	GetPublicKey([]byte)
	Agree([]byte) ([]byte, error)
}

type Keypairs struct {
	PrivateKey   x25519.PrivateKey
	KeystorePath string
	Pnt          string
}

func (k *Keypairs) Init(keystorePath string) {
	pnt := "pnt"

	k.KeystorePath = keystorePath
	k.Pnt = pnt

	privateKey, err := x25519.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}
	// k.PrivateKey = *privateKey

	var keystore = Keystore{}
	keystore.Init(keystorePath)
	keystore.Store(privateKey.Bytes(), privateKey.PublicKey.Bytes(), "pnt")
	k.PrivateKey.PublicKey.SetBytes(privateKey.PublicKey.Bytes())
}

func (k *Keypairs) GetPublicKey() {
	publicKey, err := curve25519.X25519(k.PrivateKey.Bytes(), curve25519.Basepoint)
	if err != nil {
		panic(err)
	}
	k.PrivateKey.PublicKey.SetBytes(publicKey)
}

func (k Keypairs) Agree(peerPublicKeyRaw []byte) ([]byte, error) {
	var peerPublicKey x25519.PublicKey
	peerPublicKey.SetBytes(peerPublicKeyRaw)

	var keystore = Keystore{}
	keystore.Init(k.KeystorePath)

	priKey, _, err := keystore.Fetch(k.Pnt)

	fmt.Println(priKey)

	if err != nil {
		return []byte{}, err
	}

	k.PrivateKey.SetBytes(priKey)
	return k.PrivateKey.Shared(&peerPublicKey), nil
}
