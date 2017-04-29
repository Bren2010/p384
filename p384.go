//+build ignore

package p384

import (
	"crypto/elliptic"
	"math/big"
)

type Curve struct{}

func (c *Curve) Params() *elliptic.CurveParams {
	return elliptic.P384().Params()
}

func (c *Curve) IsOnCurve(x, y *big.Int) bool {

}

func (c *Curve) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {

}

func (c *Curve) Double(x1, y1 *big.Int) (x, y *big.Int) {

}

func (c *Curve) ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) {

}

func (c *Curve) ScalarBaseMult(k []byte) (x, y *big.Int) {

}

func (c *Curve) CombinedMult(bigX, bigY *big.Int, baseScalar, scalar []byte) (x, y *big.Int) {

}
