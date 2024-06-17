//go:build !noasm && arm64
// AUTO-GENERATED BY GOAT -- DO NOT EDIT

TEXT ·dot_sve(SB), $0-32
	MOVD a+0(FP), R0
	MOVD b+8(FP), R1
	MOVD res+16(FP), R2
	MOVD len+24(FP), R3
	WORD $0xf9400068    // ldr	x8, [x3]
	WORD $0x04a0e3ea    // cntw	x10
	WORD $0xcb0a03e9    // neg	x9, x10
	WORD $0x04bf502c    // rdvl	x12, #1
	WORD $0x2598e3e0    // ptrue	p0.s
	WORD $0x8a090109    // and	x9, x8, x9
	WORD $0xeb09019f    // cmp	x12, x9
	WORD $0x540000e9    // b.ls	.LBB0_2
	WORD $0x25b8c000    // mov	z0.s, #0
	WORD $0xaa1f03eb    // mov	x11, xzr
	WORD $0x04603001    // mov	z1.d, z0.d
	WORD $0x04603002    // mov	z2.d, z0.d
	WORD $0x04603003    // mov	z3.d, z0.d
	WORD $0x14000027    // b	.LBB0_5

LBB0_2:
	WORD $0x25b8c004 // mov	z4.s, #0
	WORD $0x04bf5070 // rdvl	x16, #3
	WORD $0xaa1f03eb // mov	x11, xzr
	WORD $0x8b0c000f // add	x15, x0, x12
	WORD $0x8b0c0032 // add	x18, x1, x12
	WORD $0x04643086 // mov	z6.d, z4.d
	WORD $0x04bf5051 // rdvl	x17, #2
	WORD $0x04643087 // mov	z7.d, z4.d
	WORD $0x04643085 // mov	z5.d, z4.d
	WORD $0x8b10000d // add	x13, x0, x16
	WORD $0x8b11000e // add	x14, x0, x17
	WORD $0x8b100030 // add	x16, x1, x16
	WORD $0x8b110031 // add	x17, x1, x17

LBB0_3:
	WORD $0xa54b4003 // ld1w	{ z3.s }, p0/z, [x0, x11, lsl #2]
	WORD $0xa54b41e2 // ld1w	{ z2.s }, p0/z, [x15, x11, lsl #2]
	WORD $0xa54b41c1 // ld1w	{ z1.s }, p0/z, [x14, x11, lsl #2]
	WORD $0xa54b41a0 // ld1w	{ z0.s }, p0/z, [x13, x11, lsl #2]
	WORD $0xa54b4030 // ld1w	{ z16.s }, p0/z, [x1, x11, lsl #2]
	WORD $0x65a48203 // fmad	z3.s, p0/m, z16.s, z4.s
	WORD $0xa54b4251 // ld1w	{ z17.s }, p0/z, [x18, x11, lsl #2]
	WORD $0xa54b4232 // ld1w	{ z18.s }, p0/z, [x17, x11, lsl #2]
	WORD $0xa54b4213 // ld1w	{ z19.s }, p0/z, [x16, x11, lsl #2]
	WORD $0x65a68222 // fmad	z2.s, p0/m, z17.s, z6.s
	WORD $0x65a78241 // fmad	z1.s, p0/m, z18.s, z7.s
	WORD $0x65a58260 // fmad	z0.s, p0/m, z19.s, z5.s
	WORD $0x8b0c016b // add	x11, x11, x12
	WORD $0x8b0b0183 // add	x3, x12, x11
	WORD $0xeb09007f // cmp	x3, x9
	WORD $0x04633064 // mov	z4.d, z3.d
	WORD $0x04623046 // mov	z6.d, z2.d
	WORD $0x04613027 // mov	z7.d, z1.d
	WORD $0x04603005 // mov	z5.d, z0.d
	WORD $0x54fffda9 // b.ls	.LBB0_3
	WORD $0x14000005 // b	.LBB0_5

LBB0_4:
	WORD $0xa54b4004 // ld1w	{ z4.s }, p0/z, [x0, x11, lsl #2]
	WORD $0xa54b4025 // ld1w	{ z5.s }, p0/z, [x1, x11, lsl #2]
	WORD $0x8b0a016b // add	x11, x11, x10
	WORD $0x65a400a3 // fmla	z3.s, p0/m, z5.s, z4.s

LBB0_5:
	WORD $0xeb09017f // cmp	x11, x9
	WORD $0x54ffff63 // b.lo	.LBB0_4
	WORD $0x65802063 // faddv	s3, p0, z3.s
	WORD $0xeb08013f // cmp	x9, x8
	WORD $0x65802042 // faddv	s2, p0, z2.s
	WORD $0x1e222862 // fadd	s2, s3, s2
	WORD $0x65802021 // faddv	s1, p0, z1.s
	WORD $0x65802000 // faddv	s0, p0, z0.s
	WORD $0x1e212841 // fadd	s1, s2, s1
	WORD $0x1e202820 // fadd	s0, s1, s0
	WORD $0x54000520 // b.eq	.LBB0_13
	WORD $0xb240012a // orr	x10, x9, #0x1
	WORD $0xeb0a011f // cmp	x8, x10
	WORD $0x9a8a810a // csel	x10, x8, x10, hi
	WORD $0xcb09014b // sub	x11, x10, x9
	WORD $0x0460e3ea // cnth	x10
	WORD $0xeb0a017f // cmp	x11, x10
	WORD $0x54000062 // b.hs	.LBB0_9
	WORD $0xaa0903ea // mov	x10, x9
	WORD $0x14000019 // b	.LBB0_12

LBB0_9:
	WORD $0xa9bf7bfd // stp	x29, x30, [sp, #-16]!
	WORD $0xcb0a03ed // neg	x13, x10
	WORD $0x04bf504f // rdvl	x15, #2
	WORD $0x8b09080e // add	x14, x0, x9, lsl #2
	WORD $0x910003fd // mov	x29, sp
	WORD $0x8a0d016c // and	x12, x11, x13
	WORD $0x8b0c012a // add	x10, x9, x12
	WORD $0x8b090829 // add	x9, x1, x9, lsl #2
	WORD $0xaa0c03f0 // mov	x16, x12

LBB0_10:
	WORD $0xa540a1c1 // ld1w	{ z1.s }, p0/z, [x14]
	WORD $0xa540a123 // ld1w	{ z3.s }, p0/z, [x9]
	WORD $0xab0d0210 // adds	x16, x16, x13
	WORD $0x65830821 // fmul	z1.s, z1.s, z3.s
	WORD $0xa541a1c2 // ld1w	{ z2.s }, p0/z, [x14, #1, mul vl]
	WORD $0x8b0f01ce // add	x14, x14, x15
	WORD $0x65982020 // fadda	s0, p0, s0, z1.s
	WORD $0xa541a121 // ld1w	{ z1.s }, p0/z, [x9, #1, mul vl]
	WORD $0x8b0f0129 // add	x9, x9, x15
	WORD $0x65810841 // fmul	z1.s, z2.s, z1.s
	WORD $0x65982020 // fadda	s0, p0, s0, z1.s
	WORD $0x54fffea1 // b.ne	.LBB0_10
	WORD $0xeb0c017f // cmp	x11, x12
	WORD $0xa8c17bfd // ldp	x29, x30, [sp], #16
	WORD $0x54000100 // b.eq	.LBB0_13

LBB0_12:
	WORD $0xbc6a7801 // ldr	s1, [x0, x10, lsl #2]
	WORD $0xbc6a7822 // ldr	s2, [x1, x10, lsl #2]
	WORD $0x9100054a // add	x10, x10, #1
	WORD $0xeb08015f // cmp	x10, x8
	WORD $0x1e220821 // fmul	s1, s1, s2
	WORD $0x1e212800 // fadd	s0, s0, s1
	WORD $0x54ffff43 // b.lo	.LBB0_12

LBB0_13:
	WORD $0xbd000040 // str	s0, [x2]
	WORD $0xd65f03c0 // ret