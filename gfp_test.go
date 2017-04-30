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

func TestMul(t *testing.T) {
	P := elliptic.P384().Params().P
	R := big.NewInt(1)
	R.Lsh(R, 384).Mod(R, P).ModInverse(R, P)

	for i := 0; i < 100000; i++ {
		x, _ := rand.Int(rand.Reader, P)
		y, _ := rand.Int(rand.Reader, P)
		X, Y := &gfP{}, &gfP{}
		copy(X[:], x.Bits())
		copy(Y[:], y.Bits())

		x.Mul(x, y).Mul(x, R).Mod(x, P)
		gfpMul(X, X, Y)

		if x.Cmp(X.Int()) != 0 {
			t.Fatal("not equal")
		}
	}
}

func BenchmarkMul(b *testing.B) {
	P := elliptic.P384().Params().P
	x, _ := rand.Int(rand.Reader, P)
	X := &gfP{}
	copy(X[:], x.Bits())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gfpMul(X, X, X)
	}
}

func BenchmarkBigMul(b *testing.B) {
	P := elliptic.P384().Params().P
	x, _ := rand.Int(rand.Reader, P)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Mul(x, x).Mod(x, P)
	}
}

func BenchmarkScalarBase(b *testing.B) {
	P := elliptic.P384().Params().P
	x, _ := rand.Int(rand.Reader, P)
	X := x.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		elliptic.P384().ScalarBaseMult(X)
	}
}
