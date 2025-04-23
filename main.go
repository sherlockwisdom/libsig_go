package main

import "fmt"

func main() {
	pk, _ := GetPublicKey()
	pubKey := pk.PublicKey.Bytes()

	fmt.Println("pubkey: ", pubKey)
	fmt.Println("pubkey len: ", len(pubKey))
}
