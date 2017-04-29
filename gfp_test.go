package p384

import (
	"testing"

	"crypto/elliptic"
	"crypto/rand"
)

func TestNegZero(t *testing.T) {
	zero, x := &gfP{}, &gfP{}
	gfpNeg(x, zero)

	if *x != *zero {
		t.Fatal()
	}
}

func TestNeg(t *testing.T) {
	P := elliptic.P384().Params().P

	for i := 0; i < 20000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		X := &gfP{}
		copy(X[:], x.Bits())

		x.Neg(x).Mod(x, P)
		gfpNeg(X, X)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}

func TestAdd(t *testing.T) {
	P := elliptic.P384().Params().P

	for i := 0; i < 10000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		y, _ := rand.Int(rand.Reader, P)
		X, Y := &gfP{}, &gfP{}
		copy(X[:], x.Bits())
		copy(Y[:], y.Bits())

		x.Add(x, y).Mod(x, P)
		gfpAdd(X, X, Y)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}

func TestSub(t *testing.T) {
	P := elliptic.P384().Params().P

	for i := 0; i < 10000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		y, _ := rand.Int(rand.Reader, P)
		X, Y := &gfP{}, &gfP{}
		copy(X[:], x.Bits())
		copy(Y[:], y.Bits())

		x.Sub(x, y).Mod(x, P)
		gfpSub(X, X, Y)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}
