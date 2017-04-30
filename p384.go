package p384

import (
	"crypto/elliptic"
	"math/big"
)

var (
	// p is the order of the base field, represented as little-endian 64-bit words.
	p = gfP{0xffffffff, 0xffffffff00000000, 0xfffffffffffffffe, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff}

	// pp satisfies r*rp - p*pp = 1 where rp and pp are both integers.
	pp = gfP{0x100000001, 0x1, 0xfffffffbfffffffe, 0xfffffffcfffffffa, 0xc00000002, 0x1400000014}

	// r2 is R^2 where R = 2^384 mod p.
	r2 = gfP{0xfffffffe00000001, 0x200000000, 0xfffffffe00000000, 0x200000000, 0x1}

	// r3 is R^3 where R = 2^384 mod p.
	r3 = gfP{0xfffffffc00000002, 0x300000002, 0xfffffffcfffffffe, 0x300000005, 0xfffffffdfffffffd, 0x300000002}

	// rN1 is R^-1 where R = 2^384 mod p.
	rN1 = gfP{0xffffffe100000006, 0xffffffebffffffd8, 0xfffffffbfffffffd, 0xfffffffcfffffffa, 0xc00000002, 0x1400000014}

	// b is the curve's B parameter, Montgomery encoded.
	b = gfP{0x81188719d412dcc, 0xf729add87a4c32ec, 0x77f2209b1920022e, 0xe3374bee94938ae2, 0xb62b21f41f022094, 0xcd08114b604fbff9}
)

type Curve struct{}

func (c *Curve) Params() *elliptic.CurveParams {
	return elliptic.P384().Params()
}

func (c *Curve) IsOnCurve(X, Y *big.Int) bool {
	x, y := &gfP{}, &gfP{}
	copy(x[:], X.Bits())
	copy(y[:], Y.Bits())
	montEncode(x, x)
	montEncode(y, y)

	y2, x3 := &gfP{}, &gfP{}
	gfpMul(y2, y, y)
	gfpMul(x3, x, x)
	gfpMul(x3, x3, x)

	threeX := &gfP{}
	gfpAdd(threeX, x, x)
	gfpAdd(threeX, threeX, x)

	gfpSub(x3, x3, threeX)
	gfpAdd(x3, x3, &b)

	return *y2 == *x3
}

// func (c *Curve) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
//
// }
//
// func (c *Curve) Double(x1, y1 *big.Int) (x, y *big.Int) {
//
// }
//
// func (c *Curve) ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) {
//
// }
//
// func (c *Curve) ScalarBaseMult(k []byte) (x, y *big.Int) {
//
// }
//
// func (c *Curve) CombinedMult(bigX, bigY *big.Int, baseScalar, scalar []byte) (x, y *big.Int) {
//
// }
