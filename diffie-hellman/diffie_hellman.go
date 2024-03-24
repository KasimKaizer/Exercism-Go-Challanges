// Package diffiehellman contains various tools to implement Diffie-Hellman key exchange
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey creates a new private key, it achieves this generating a random big.int which is
// less then the given n but greater than 1.
func PrivateKey(n *big.Int) *big.Int {
	rand, _ := rand.Int(rand.Reader, n)
	if rand.Cmp(big.NewInt(1)) != 1 {
		return PrivateKey(n)
	}
	return rand
}

// PublicKey creates a new public key, it does this by using this formula:
// public key = g^private MOD n
func PublicKey(private, n *big.Int, g int64) *big.Int {
	bigG := big.NewInt(g)
	return new(big.Int).Exp(bigG, private, n)
}

// NewPair creates a new pair of a private and a public key.
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	priKey := PrivateKey(p)
	pubKey := PublicKey(priKey, p, g)
	return priKey, pubKey
}

// SecretKey evaluates the secret key from the given values and returns it, it uses the formula:
// secret = public2^private1 MOD n
func SecretKey(private1, public2, n *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, n)
}
