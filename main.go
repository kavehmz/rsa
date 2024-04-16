package main

import (
	"fmt"
	"math/big"
)

func main() {
	priv := private{p: big.NewInt(400009), q: big.NewInt(500009), n: big.NewInt(0)}
	priv.n.Mul(priv.p, priv.q)

	pM1 := big.NewInt(0).Sub(priv.p, big.NewInt(1))
	qM1 := big.NewInt(0).Sub(priv.q, big.NewInt(1))
	priv.λ = lcm(pM1, qM1)
	priv.e = big.NewInt(65537)
	priv.d = big.NewInt(0).ModInverse(priv.e, priv.λ)

	fmt.Println("p, q, n:", priv.p, priv.q, priv.n)
	fmt.Println("λ(n):", priv.λ)
	fmt.Println("e (default):", 65537)
	fmt.Println("d:", priv.d)

	m := big.NewInt(65)
	pub := public{n: priv.n, e: priv.e}
	// encrypt
	c := pub.enc(m)
	fmt.Println(c)
	// decrypt
	dc := priv.dec(c)

	fmt.Println("value to encrypt:", big.NewInt(65))
	fmt.Println("encrypted:", c)
	fmt.Println("decrypted:", dc)

}

type private struct {
	p *big.Int // prime #1
	q *big.Int // prime #2
	n *big.Int // p*q
	λ *big.Int // Carmichael's totient function
	e *big.Int // co_prime to (p-1)*(q-1)
	d *big.Int
}

type public struct {
	n *big.Int // p*q
	e *big.Int // co_prime to (p-1)*(q-1)
}

func (s *public) enc(m *big.Int) *big.Int {
	return big.NewInt(0).Exp(m, s.e, s.n)
}

func (s *private) dec(c *big.Int) *big.Int {
	return big.NewInt(0).Exp(c, s.d, s.n)
}

func lcm(a, b *big.Int) *big.Int {

	mul := big.NewInt(1)
	gcd := big.NewInt(1)
	mul.Mul(a, b)
	gcd.GCD(nil, nil, a, b)
	return mul.Div(mul, gcd)
}
