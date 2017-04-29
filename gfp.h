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

#define mul(ra, rb) \
	mulArb(0+ra,8+ra,16+ra,24+ra,32+ra,40+ra, rb)

#define mulArb(a1,a2,a3,a4,a5,a6, rb) \
	MOVQ a1, DX \
	MULXQ 0+rb, R8, R9 \
	MULXQ 8+rb, AX, R10 \
	ADDQ AX, R9 \
	MULXQ 16+rb, AX, R11 \
	ADCQ AX, R10 \
	MULXQ 24+rb, AX, R12 \
	ADCQ AX, R11 \
	MULXQ 32+rb, AX, R13 \
	ADCQ AX, R12 \
	MULXQ 40+rb, AX, R14 \
	ADCQ AX, R13 \
	ADCQ $0, R14 \
	\
	MOVQ R8, 0(SP) \
	MOVQ $0, R15 \
	MOVQ $0, R8 \
	\
	MOVQ a2, DX \
	MULXQ 0+rb, AX, BX \
	ADDQ AX, R9 \
	ADCQ BX, R10 \
	MULXQ 16+rb, AX, BX \
	ADCQ AX, R11 \
	ADCQ BX, R12 \
	MULXQ 32+rb, AX, BX \
	ADCQ AX, R13 \
	ADCQ BX, R14 \
	ADCQ $0, R15 \
	MULXQ 8+rb, AX, BX \
	ADDQ AX, R10 \
	ADCQ BX, R11 \
	MULXQ 24+rb, AX, BX \
	ADCQ AX, R12 \
	ADCQ BX, R13 \
	MULXQ 40+rb, AX, BX \
	ADCQ AX, R14 \
	ADCQ BX, R15 \
	ADCQ $0, R8 \
	\
	MOVQ R9, 8(SP) \
	MOVQ $0, R9 \
	\
	MOVQ a3, DX \
	MULXQ 0+rb, AX, BX \
	ADDQ AX, R10 \
	ADCQ BX, R11 \
	MULXQ 16+rb, AX, BX \
	ADCQ AX, R12 \
	ADCQ BX, R13 \
	MULXQ 32+rb, AX, BX \
	ADCQ AX, R14 \
	ADCQ BX, R15 \
	ADCQ $0, R8 \
	MULXQ 8+rb, AX, BX \
	ADDQ AX, R11 \
	ADCQ BX, R12 \
	MULXQ 24+rb, AX, BX \
	ADCQ AX, R13 \
	ADCQ BX, R14 \
	MULXQ 40+rb, AX, BX \
	ADCQ AX, R15 \
	ADCQ BX, R8 \
	ADCQ $0, R9 \
	\
	MOVQ R10, 16(SP) \
	MOVQ $0, R10 \
	\
	MOVQ a4, DX \
	MULXQ 0+rb, AX, BX \
	ADDQ AX, R11 \
	ADCQ BX, R12 \
	MULXQ 16+rb, AX, BX \
	ADCQ AX, R13 \
	ADCQ BX, R14 \
	MULXQ 32+rb, AX, BX \
	ADCQ AX, R15 \
	ADCQ BX, R8 \
	ADCQ $0, R9 \
	MULXQ 8+rb, AX, BX \
	ADDQ AX, R12 \
	ADCQ BX, R13 \
	MULXQ 24+rb, AX, BX \
	ADCQ AX, R14 \
	ADCQ BX, R15 \
	MULXQ 40+rb, AX, BX \
	ADCQ AX, R8 \
	ADCQ BX, R9 \
	ADCQ $0, R10 \
	\
	MOVQ R11, 24(SP) \
	MOVQ $0, R11 \
	\
	MOVQ a5, DX \
	MULXQ 0+rb, AX, BX \
	ADDQ AX, R12 \
	ADCQ BX, R13 \
	MULXQ 16+rb, AX, BX \
	ADCQ AX, R14 \
	ADCQ BX, R15 \
	MULXQ 32+rb, AX, BX \
	ADCQ AX, R8 \
	ADCQ BX, R9 \
	ADCQ $0, R10 \
	MULXQ 8+rb, AX, BX \
	ADDQ AX, R13 \
	ADCQ BX, R14 \
	MULXQ 24+rb, AX, BX \
	ADCQ AX, R15 \
	ADCQ BX, R8 \
	MULXQ 40+rb, AX, BX \
	ADCQ AX, R9 \
	ADCQ BX, R10 \
	ADCQ $0, R11 \
	\
	MOVQ R12, 32(SP) \
	\
	MOVQ a6, DX \
	MULXQ 0+rb, AX, BX \
	ADDQ AX, R13 \
	ADCQ BX, R14 \
	MULXQ 16+rb, AX, BX \
	ADCQ AX, R15 \
	ADCQ BX, R8 \
	MULXQ 32+rb, AX, BX \
	ADCQ AX, R9 \
	ADCQ BX, R10 \
	ADCQ $0, R11 \
	MULXQ 8+rb, AX, BX \
	ADDQ AX, R14 \
	ADCQ BX, R15 \
	MULXQ 24+rb, AX, BX \
	ADCQ AX, R8 \
	ADCQ BX, R9 \
	MULXQ 40+rb, AX, BX \
	ADCQ AX, R10 \
	ADCQ BX, R11 \
	\
	MOVQ R13, 40(SP) \
	MOVQ R14, 48(SP) \
	MOVQ R15, 56(SP) \
	MOVQ R8,  64(SP) \
	MOVQ R9,  72(SP) \
	MOVQ R10, 80(SP) \
	MOVQ R11, 88(SP)

#define gfpReduce() \
	\ // m = (T * N') mod R, store m in R8:R9:R10:R11
	MOVQ ·np+0(SB), DX \
	MULXQ 0(SP), R8, R9 \
	MULXQ 8(SP), AX, R10 \
	ADDQ AX, R9 \
	MULXQ 16(SP), AX, R11 \
	ADCQ AX, R10 \
	MULXQ 24(SP), AX, BX \
	ADCQ AX, R11 \
	\
	MOVQ ·np+8(SB), DX \
	MULXQ 0(SP), AX, BX \
	ADDQ AX, R9 \
	ADCQ BX, R10 \
	MULXQ 16(SP), AX, BX \
	ADCQ AX, R11 \
	MULXQ 8(SP), AX, BX \
	ADDQ AX, R10 \
	ADCQ BX, R11 \
	\
	MOVQ ·np+16(SB), DX \
	MULXQ 0(SP), AX, BX \
	ADDQ AX, R10 \
	ADCQ BX, R11 \
	MULXQ 8(SP), AX, BX \
	ADDQ AX, R11 \
	\
	MOVQ ·np+24(SB), DX \
	MULXQ 0(SP), AX, BX \
	ADDQ AX, R11 \
	\
	storeBlock(R8,R9,R10,R11, 64(SP)) \
	\
	\ // m * N
	mulArb(·p2+0(SB),·p2+8(SB),·p2+16(SB),·p2+24(SB), 64(SP)) \
	\
	\ // Add the 512-bit intermediate to m*N
	MOVQ $0, AX \
	ADDQ 0(SP), R8 \
	ADCQ 8(SP), R9 \
	ADCQ 16(SP), R10 \
	ADCQ 24(SP), R11 \
	ADCQ 32(SP), R12 \
	ADCQ 40(SP), R13 \
	ADCQ 48(SP), R14 \
	ADCQ 56(SP), R15 \
	ADCQ $0, AX \
	\
	gfpCarry(R12,R13,R14,R15,AX, R8,R9,R10,R11,BX)
