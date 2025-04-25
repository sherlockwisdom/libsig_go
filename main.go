package main

import "fmt"

func main() {
	pk, _ := Init()

	fmt.Println("pubkey: ", pk.Bytes())
	fmt.Println("pubkey len: ", len(pk.Bytes()))
}
