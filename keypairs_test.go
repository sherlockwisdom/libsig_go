package main

import (
	"fmt"
	"testing"
	"reflect"
	"bytes"
)

func TestSK(t *testing.T) {
	pk, prKBytes, _ := Init()
	pk1, prK1Bytes, _ := Init()

	fmt.Println(reflect.TypeOf(pk.Bytes()))

	if pk == pk1 {
		t.Errorf("PK are equal: got %d, %d", pk.Bytes(), pk1.Bytes())
	}

	sk := Agree(pk1.Bytes(), prKBytes)
	sk1 := Agree(pk.Bytes(), prK1Bytes)

	/*
	if reflect.DeepEqual(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
	*/

	if bytes.Equal(sk, sk1) == false {
		t.Errorf("SK not equal: got %s, want: %s", sk, sk1)
	}
}
