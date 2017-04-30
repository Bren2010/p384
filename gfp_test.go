package p384

import (
	"testing"

	"crypto/elliptic"
	"crypto/rand"
	"math/big"
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

func TestMulZero(t *testing.T) {
	P := elliptic.P384().Params().P
	x, _ := rand.Int(rand.Reader, P)
	X := &gfP{}
	copy(X[:], x.Bits())

	zero := &gfP{}
	gfpMul(X, X, zero)

	if *X != *zero {
		t.Fatal("not zero")
	}
}

func TestMul(t *testing.T) {
	P := elliptic.P384().Params().P
	Rinv := big.NewInt(1)
	Rinv.Lsh(Rinv, 384).Mod(Rinv, P).ModInverse(Rinv, P)

	for i := 0; i < 10000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		y, _ := rand.Int(rand.Reader, P)
		X, Y := &gfP{}, &gfP{}
		copy(X[:], x.Bits())
		copy(Y[:], y.Bits())

		x.Mul(x, y).Mul(x, Rinv).Mod(x, P)
		gfpMul(X, X, Y)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}

func TestInvert(t *testing.T) {
	P := elliptic.P384().Params().P

	for i := 0; i < 1000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		X := &gfP{}
		copy(X[:], x.Bits())

		x.ModInverse(x, P)
		montEncode(X, X)
		X.Invert(X)
		montDecode(X, X)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}
