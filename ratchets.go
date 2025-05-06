package main

import (
	"io"
	"crypto/sha512"
	"crypto/sha256"
	"golang.org/x/crypto/hkdf"
	"crypto/hmac"
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

func BobInit(
	state *States,
	SK []byte,
	bobPrivateKey []byte,
) {
	state.DHs = bobPrivateKey
        state.DHr = nil
        state.RK = SK
        state.CKs = nil
        state.CKr = nil
        state.Ns = 0
        state.Nr = 0
        state.PN = 0
	state.MKSKIPPED = make(map[string]int)
}

func GENERATE_DH() []byte {
	pk := KeypairInit()
	return pk.Bytes()
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

func DH(privKey []byte, pubKey []byte) []byte {
	return KeypairAgree(privKey, pubKey)
}

func Encrypt(state *States, data []byte, AD []byte) (Headers, []byte) {
	var mk = make([]byte, 0)
	state.CKs, mk = KDF_CK(state.CKs)
	header := Headers {
		state.DHs,
		state.PN,
		state.Ns,
	}
	state.Ns += 1

	// TODO:
	// return header, ENCRYPT(mk, data, CONCAT(AD, header))
	// return header, []byte{}
	return header, mk
}

func KDF_CK(ck []byte) ([]byte, []byte) {
	tck := hmac.New(sha256.New, ck)
	CK := tck.Sum([]byte{0x01})
	MK := tck.Sum([]byte{0x02})

	return CK, MK
}

/*
func CONCAT(ad []byte, header Headers) []byte {
	serializedHeader := SerializeProtocols(header)
	result := make([]byte, len(ad) + len(serializedHeader))
	copy(result, ad)
	copy(result[len(a):], serializedHeader)
	return result
}
*/

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

func Decrypt() []byte {
}

*/
