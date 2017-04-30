package p384

import (
	"testing"

	"crypto/elliptic"
	"crypto/rand"
)

func TestIsOnCurveTrue(t *testing.T) {
	for i := 0; i < 100; i++ {
		K := make([]byte, 384/8)
		rand.Read(K)

		X, Y := elliptic.P384().ScalarBaseMult(K)

		c := &Curve{}
		if !c.IsOnCurve(X, Y) {
			t.Fatal("not on curve")
		}
	}
}

func TestIsOnCurveFalse(t *testing.T) {
	P := elliptic.P384().Params().P

	for i := 0; i < 100; i++ {
		X, _ := rand.Int(rand.Reader, P)
		Y, _ := rand.Int(rand.Reader, P)

		c := &Curve{}
		if c.IsOnCurve(X, Y) {
			t.Fatal("bad point on curve")
		}
	}
}
