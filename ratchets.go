package main

import (
	"io"
	"crypto/sha512"
	"golang.org/x/crypto/hkdf"
	"maze.io/x/crypto/x25519"
)

func AliceInit(
	state *States, 
	SK []byte, 
	bobPubKey []byte,
) {
	state.DHs = GENERATE_DH()
	state.DHr = bobPubKey
	state.RK, state.CKs = KDF_RK(SK, DH(state.DHs, state.DHr))
	state.CKr = nil
	state.Ns = 0
	state.Nr = 0
	state.PN = 0
	state.MKSKIPPED = make(map[string]int)
}

func BobInit() {
}

func GENERATE_DH() x25519.PrivateKey {
	pk := KeypairInit()
	return pk
}

func KDF_RK(RK, DHOut []byte) ([]byte, []byte) {
	Len := 32

	info := []byte("KDF_RK")

	hash := sha512.New
	hkdf := hkdf.New(hash, DHOut, RK, info)

	rk, ck := make([]byte, Len), make([]byte, Len)

	if _, err := io.ReadFull(hkdf, rk); err != nil {
		panic(err)
	}

	if _, err := io.ReadFull(hkdf, ck); err != nil {
		panic(err)
	}
	return rk, ck
}

func DH(privKey x25519.PrivateKey, pubKey []byte) []byte {
	return KeypairAgree(privKey, pubKey)
}

/*
func DHRatchet(state States, header Headers) {
	state.PN = state.Ns
	state.Ns = 0
	state.Nr = 0
	state.DHr = header.Dh
	state.RK, state.CKr = KDF_RK(state.Rk, DH(state.DHs, state.DHr))
	state.DHs = GENERATE_DH()
	state.RK, state.CKs = KDF_RK(state.Rk, DH(state.DHs, state.DHr))
}



func KDF_CK() []byte {
}

func Encrypt() []byte {
}

func Decrypt() []byte {
}

func CONCAT() []byte {
}
*/
