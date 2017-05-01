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

func (c *Curve) add(a *jacobianPoint, b *affinePoint) *jacobianPoint {
	if a.IsZero() {
		return b.ToJacobian()
	} else if b.IsZero() {
		return a.Dup()
	}

	z1z1, u2 := &gfP{}, &gfP{}
	gfpMul(z1z1, &a.z, &a.z)
	gfpMul(u2, &b.x, z1z1)

	s2 := &gfP{}
	gfpMul(s2, &b.y, &a.z)
	gfpMul(s2, s2, z1z1)
	if a.x == *u2 {
		if a.y != *s2 {
			return &jacobianPoint{}
		}
		return c.double(a)
	}

	h, r := &gfP{}, &gfP{}
	gfpSub(h, u2, &a.x)
	gfpSub(r, s2, &a.y)

	h2, h3 := &gfP{}, &gfP{}
	gfpMul(h2, h, h)
	gfpMul(h3, h2, h)

	h2x1 := &gfP{}
	gfpMul(h2x1, h2, &a.x)

	x3, y3, z3 := &gfP{}, &gfP{}, &gfP{}
	gfpMul(x3, r, r)
	gfpSub(x3, x3, h3)
	gfpSub(x3, x3, h2x1)
	gfpSub(x3, x3, h2x1)

	gfpSub(y3, h2x1, x3)
	gfpMul(y3, y3, r)
	h3y1 := &gfP{}
	gfpMul(h3y1, h3, &a.y)
	gfpSub(y3, y3, h3y1)

	gfpMul(z3, h, &a.z)

	return &jacobianPoint{*x3, *y3, *z3}
}

func (c *Curve) double(a *jacobianPoint) *jacobianPoint {
	delta, gamma, alpha, alpha2 := &gfP{}, &gfP{}, &gfP{}, &gfP{}
	gfpMul(delta, &a.z, &a.z)
	gfpMul(gamma, &a.y, &a.y)
	gfpSub(alpha, &a.x, delta)
	gfpAdd(alpha2, &a.x, delta)
	gfpMul(alpha, alpha, alpha2)
	*alpha2 = *alpha
	gfpAdd(alpha, alpha, alpha)
	gfpAdd(alpha, alpha, alpha2)

	beta := &gfP{}
	gfpMul(beta, &a.x, gamma)

	x3, beta8 := &gfP{}, &gfP{}
	gfpMul(x3, alpha, alpha)
	gfpAdd(beta8, beta, beta)
	gfpAdd(beta8, beta8, beta8)
	gfpAdd(beta8, beta8, beta8)
	gfpSub(x3, x3, beta8)

	z3 := &gfP{}
	gfpAdd(z3, &a.y, &a.z)
	gfpMul(z3, z3, z3)
	gfpSub(z3, z3, gamma)
	gfpSub(z3, z3, delta)

	gfpAdd(beta, beta, beta)
	gfpAdd(beta, beta, beta)
	gfpSub(beta, beta, x3)

	y3 := &gfP{}
	gfpMul(y3, alpha, beta)

	gfpMul(gamma, gamma, gamma)
	gfpAdd(gamma, gamma, gamma)
	gfpAdd(gamma, gamma, gamma)
	gfpAdd(gamma, gamma, gamma)
	gfpSub(y3, y3, gamma)

	return &jacobianPoint{*x3, *y3, *z3}
}

func (c *Curve) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
	pt := c.add(newAffinePoint(x1, y1).ToJacobian(), newAffinePoint(x2, y2))
	return pt.ToAffine().ToInt()
}

func (c *Curve) Double(x1, y1 *big.Int) (x, y *big.Int) {
	pt := c.double(newAffinePoint(x1, y1).ToJacobian())
	return pt.ToAffine().ToInt()
}

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
