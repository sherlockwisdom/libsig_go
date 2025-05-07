package main

/*
	"crypto/rand"
	"maze.io/x/crypto/x25519"
	"golang.org/x/crypto/curve25519"
*/

/*
func TestDeterministicPublicKeys(t *testing.T) {
	var keypair Keypairs
	keypair.Init("db/test_deterministic.db")
	value := keypair.PrivateKey.Bytes()

	var keypair1 Keypairs
	keypair1.PrivateKey.SetBytes(value)
	keypair1.GetPublicKey()

	if !bytes.Equal(
		keypair1.PrivateKey.PublicKey.Bytes(),
		keypair.PrivateKey.PublicKey.Bytes(),
	) {
		t.Errorf("Pub keys differ! \n\twanted %d, \n\tgot %d", value, keypair1.PrivateKey.PublicKey.Bytes())
	}
}
*/
