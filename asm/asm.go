package asm

const (
	CondEQ = 0   // z set
	CondNE = 1   // z clear
	CondCS = 2   // c set
	CondCC = 3   // c clear
	CondMI = 4   // n set, minus
	CondPL = 5   // n clear, plus
	CondVS = 6   // v set, overflow
	CondVC = 7   // v clear, no overflow
	CondHI = 8   // c set and z clear,
	CondLS = 9   // c clear or z set
	CondGE = 0xA // n == v
	CondLT = 0xB // n != v
	CondGT = 0xC // z == 0, n == v
	CondLE = 0xD // z == 1 or n != v
	CondAL = 0xE // always
	CondIN = 0xF // use for additional instruction coding
)

const (
	condShift = 28
	bitMask24 = uint32((0x1 << 24) - 1)
)

func setCond(i uint32, cond uint32) uint32 {
	return (i & ^uint32(0xf0000000)) | ((cond & 0xf) << 28)
}

func Cond(i uint32) uint8 {
	return uint8(i >> 28)
}

func branch(cond, bitL, im24 uint32) uint32 {
	ret := cond << condShift
	ret |= (bitL & 0x1) << 24
	ret |= im24 & bitMask24
	return ret
}

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
func shiftIm(ro, im uint32) uint32 {
	return ((ro & 0xf) << 8) | (im & 0xff)
}

// shiftReg creates the shift-operand that uses registers
func shiftReg(mode, sh, rs, rm uint32) uint32 {
	if mode == ShReg {
		sh = 0
	} else if mode == ShRotCarry {
		sh = 0
	}

	ret := (mode & 0x7) << 4
	ret |= rm & 0xf

	if (mode & 0x1) == 0 {
		ret |= (rs & 0xf) << 8
	} else {
		ret |= (sh & 0x1f) << 7
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
// bitI is set when the sh field will be interpreted as an immediate
// bitS is set when the flags will be affected
// rn is the first input register
// rd is the output result register
func arith(cond, op, bitI, bitS, rn, rd, sh uint32) uint32 {
	ret := cond << condShift

	switch op {
	case OpTst, OpTeq, OpCmp, OpCmn:
		bitS = 1
		rd = 0
	case OpMov, OpMvn:
		rn = 0
	}

	ret |= (bitI & 0x1) << 25
	ret |= (bitS & 0x1) << 20
	ret |= (op & 0xf) << 21
	ret |= (rn & 0xf) << 16
	ret |= (rd & 0xf) << 12
	ret |= sh & 0xfff

	return ret
}

// mutiply
func mul(cond, bitS, rd, rs, rm uint32) uint32 {
	ret := cond << condShift
	ret |= 0x9 << 4
	ret |= (bitS & 0x1) << 20
	ret |= (rd & 0xf) << 16
	ret |= (rs & 0xf) << 8
	ret |= rm & 0xf
	return ret
}

const (
	pwMask     = 0x9
	pwOffset   = 0x8 // use Rn+offset, do not write back
	pwPre      = 0x9 // use Rn+offset, and write back
	pwPost     = 0x0 // use Rn, write back Rn+offset
	pwPostUser = 0x1 // use Rn, write back Rn+offset, as if in user mode
)

func addrImm(im uint32) uint32 {
	return im & 0xfff
}

const (
	AddrSll = 0       // shift right logic
	AddrSrl = 1       // shift right logic
	AddrSra = 4       // shift right arithmetic
	AddrSrr = 6       // shift right rotate
	AddrSrc = 6 + 0xf // shift right carry
)

func addrReg(mode, im, rm uint32) uint32 {
	ret := uint32(0x1 << 24) // set the I bit first
	ret |= (mode & 0x7) << 4
	if mode != AddrSrc {
		ret |= (im & 0x1f) << 7
	}
	ret |= rm & 0xf
	return ret
}
