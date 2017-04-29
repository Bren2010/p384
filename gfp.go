package p384

import (
	"fmt"
	"math/big"
)

type gfP [6]big.Word

func newGFp(x int64) (out *gfP) {
	if x >= 0 {
		out = &gfP{big.Word(x)}
	} else {
		out = &gfP{big.Word(-x)}
		gfpNeg(out, out)
	}

	// montEncode(out, out)
	return out
}

func (e *gfP) Set(f *gfP) {
	e[0] = f[0]
	e[1] = f[1]
	e[2] = f[2]
	e[3] = f[3]
	e[4] = f[4]
	e[5] = f[5]
}

func (e *gfP) Int() *big.Int {
	return new(big.Int).SetBits(e[:])
}

func (e *gfP) String() string {
	return fmt.Sprintf("%16.16x%16.16x%16.16x%16.16x%16.16x%16.16x", e[5], e[4], e[3], e[2], e[1], e[0])
}

// func (e *gfP) Invert(f *gfP) {
// 	bits := [4]uint64{0x185cac6c5e089665, 0xee5b88d120b5b59e, 0xaa6fecb86184dc21, 0x8fb501e34aa387f9}
//
// 	sum, power := &gfP{}, &gfP{}
// 	sum.Set(rN1)
// 	power.Set(f)
//
// 	for word := 0; word < 4; word++ {
// 		for bit := uint(0); bit < 64; bit++ {
// 			if (bits[word]>>bit)&1 == 1 {
// 				gfpMul(sum, sum, power)
// 			}
// 			gfpMul(power, power, power)
// 		}
// 	}
//
// 	gfpMul(sum, sum, r3)
// 	e.Set(sum)
// }

// go:noescape
func gfpNeg(c, a *gfP)

//go:noescape
func gfpAdd(c, a, b *gfP)

//go:noescape
func gfpSub(c, a, b *gfP)

//go:noescape
func gfpMul(c *[12]big.Word, a, b *gfP)
