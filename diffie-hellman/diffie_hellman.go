package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.
// well to being with let me copy and past the description in the code itself
// Step 0
// The test program supplies prime numbers p and g.
// Step 1
// Alice picks a private key, a, greater than 1 and less than p. Bob does the same to pick a private key b.
// Step 2
// Alice calculates a public key A.

func PrivateKey(n *big.Int) *big.Int {
	rand, _ := rand.Int(rand.Reader, n)
	if rand.Cmp(big.NewInt(1)) != 1 {
		return PrivateKey(n)
	}
	return rand
}

func PublicKey(private, n *big.Int, g int64) *big.Int {
	bigG := big.NewInt(g)
	return new(big.Int).Exp(bigG, private, n)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priKey := PrivateKey(p)
	pubKey := PublicKey(priKey, p, g)
	return priKey, pubKey
}

func SecretKey(private1, public2, n *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, n)
}
