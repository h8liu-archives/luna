package arm

const (
	ShSLLImm = 0
	ShSLLReg = 1
	ShSRLImm = 2
	ShSRLReg = 3
	ShSRAImm = 4
	ShSRAReg = 5
	ShRotImm = 6 // for 0, rotate carry
	ShRotReg = 7

	// not real shift bits, but some special shift modes
	ShReg      = 0x10
	ShRotCarry = 0x16
)

// shiftIm creates the shift-operand with immediate
// ro is the ammount to rotate right
func ShiftIm(im, ro uint32) uint32 {
	return ((ro & 0xf) << 8) | (im & 0xff) | (0x1 << 25)
}

// shiftReg creates the shift-operand that uses registers
// sh/rs is the amount to shift
// rm is the register
func ShiftReg(mode, rm, sh, rs uint32) uint32 {
	if mode == ShReg || mode == ShRotCarry {
		sh = 0
	}

	ret := (mode & 0x7) << 4
	ret |= rm & 0xf

	if (mode & 0x1) == 0 {
		ret |= (sh & 0x1f) << 7
	} else {
		ret |= (rs & 0xf) << 8
	}

	return ret
}

const (
	OpAnd = 0
	OpXor = 1
	OpSub = 2
	OpRsb = 3
	OpAdd = 4

	OpTst = 8
	OpTeq = 9
	OpCmp = 0xA
	OpCmn = 0xB
	OpOrr = 0xC // bit or
	OpMov = 0xD
	OpBic = 0xE // bit clear
	OpMvn = 0xF // move negative
)

// arith makes an arithmetic instruction
// cond specifies when the instruction is executed
// op is a 4-bit field that tells the operation
// bitS is set when the flags will be affected
// rn is the first input register
// rd is the output result register
func Arith(cond, op, bitS, rn, rd, sh uint32) uint32 {
	ret := cond << condShift

	switch op {
	case OpTst, OpTeq, OpCmp, OpCmn:
		bitS = 1
		rd = 0
	case OpMov, OpMvn:
		rn = 0
	}

	ret |= (bitS & 0x1) << 20
	ret |= (op & 0xf) << 21
	ret |= (rn & 0xf) << 16
	ret |= (rd & 0xf) << 12
	ret |= sh & (0xfff | (0x1 << 25))

	return ret
}

// mutiply
func Mul(cond, bitS, rd, rs, rm uint32) uint32 {
	ret := cond << condShift
	ret |= 0x9 << 4
	ret |= (bitS & 0x1) << 20
	ret |= (rd & 0xf) << 16
	ret |= (rs & 0xf) << 8
	ret |= rm & 0xf
	return ret
}
