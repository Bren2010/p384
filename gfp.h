#define storeBlock(a1,a2,a3,a4,a5,a6, r) \
	MOVQ a1,  0+r \
	MOVQ a2,  8+r \
	MOVQ a3, 16+r \
	MOVQ a4, 24+r \
	MOVQ a5, 32+r \
	MOVQ a6, 40+r \

#define loadBlock(r, a1,a2,a3,a4,a5,a6) \
	MOVQ  0+r, a1 \
	MOVQ  8+r, a2 \
	MOVQ 16+r, a3 \
	MOVQ 24+r, a4 \
	MOVQ 32+r, a5 \
	MOVQ 40+r, a6

#define gfpCarry(a1,a2,a3,a4,a5,a6,a7, b1,b2,b3,b4,b5,b6,b7) \
	\ // b = a-p
	MOVQ a1, b1 \
	MOVQ a2, b2 \
	MOVQ a3, b3 \
	MOVQ a4, b4 \
	MOVQ a5, b5 \
	MOVQ a6, b6 \
	MOVQ a7, b7 \
	\
	SUBQ ·p+0(SB), b1 \
	SBBQ ·p+8(SB), b2 \
	SBBQ ·p+16(SB), b3 \
	SBBQ ·p+24(SB), b4 \
	SBBQ ·p+32(SB), b5 \
	SBBQ ·p+40(SB), b6 \
	SBBQ $0, b7 \
	\
	\ // if b is negative then return a
	\ // else return b
	CMOVQCC b1, a1 \
	CMOVQCC b2, a2 \
	CMOVQCC b3, a3 \
	CMOVQCC b4, a4 \
	CMOVQCC b5, a5 \
	CMOVQCC b6, a6
