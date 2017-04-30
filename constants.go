package p384

// p is the order of the base field, represented as little-endian 64-bit words.
var p = gfP{0xffffffff, 0xffffffff00000000, 0xfffffffffffffffe, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff}

// pp satisfies r*rp - p*pp = 1 where rp and pp are both integers.
var pp = gfP{0x100000001, 0x1, 0xfffffffbfffffffe, 0xfffffffcfffffffa, 0xc00000002, 0x1400000014}

// r2 is R^2 where R = 2^384 mod p.
var r2 = gfP{0xfffffffe00000001, 0x200000000, 0xfffffffe00000000, 0x200000000, 0x1}

// r3 is R^3 where R = 2^384 mod p.
var r3 = gfP{0xfffffffc00000002, 0x300000002, 0xfffffffcfffffffe, 0x300000005, 0xfffffffdfffffffd, 0x300000002}

// rN1 is R^-1 where R = 2^384 mod p.
var rN1 = gfP{0xffffffe100000006, 0xffffffebffffffd8, 0xfffffffbfffffffd, 0xfffffffcfffffffa, 0xc00000002, 0x1400000014}
