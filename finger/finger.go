package finger

import (
	. "github.com/h8liu/luna/arm"
)

// register only operations

func r3(op, ret, r1, r2 uint32) uint32 {
	return Arith(CondAL, op, 0, r1, ret, ShiftReg(ShSLLImm, r2, 0, 0))
}

func And(ret, r1, r2 uint32) uint32 { return r3(OpAnd, ret, r1, r2) }
func Xor(ret, r1, r2 uint32) uint32 { return r3(OpXor, ret, r1, r2) }
func Sub(ret, r1, r2 uint32) uint32 { return r3(OpSub, ret, r1, r2) }
func Cmp(r1, r2 uint32) uint32      { return r3(OpCmp, 0, r1, r2) }
func Add(ret, r1, r2 uint32) uint32 { return r3(OpAdd, ret, r1, r2) }
func Or(ret, r1, r2 uint32) uint32  { return r3(OpOrr, ret, r1, r2) }
func Mov(ret, r1 uint32) uint32     { return r3(OpMov, ret, 0, r1) }

// register and 8-bit

func r3i(op, ret, r1, im uint32) uint32 {
	return Arith(CondAL, op, 0, r1, ret, ShiftIm(im, 0))
}

func Andi(ret, r1, im uint32) uint32 { return r3i(OpAnd, ret, r1, im) }
func Xori(ret, r1, im uint32) uint32 { return r3i(OpXor, ret, r1, im) }
func Subi(ret, r1, im uint32) uint32 { return r3i(OpSub, ret, r1, im) }
func Addi(ret, r1, im uint32) uint32 { return r3i(OpAdd, ret, r1, im) }
func Ori(ret, r1, im uint32) uint32  { return r3i(OpOrr, ret, r1, im) }
func Movi(ret, im uint32) uint32     { return r3i(OpMov, ret, 0, im) }

// branch operations

func Beq(im int32) uint32 { return Branch(CondEQ, 0, BranchOffset(im)) }
func Bne(im int32) uint32 { return Branch(CondNE, 0, BranchOffset(im)) }
func Bge(im int32) uint32 { return Branch(CondMI, 0, BranchOffset(im)) }
func Bl(im int32) uint32  { return Branch(CondPL, 0, BranchOffset(im)) }
func J(im int32) uint32   { return Branch(CondAL, 0, BranchOffset(im)) }
func Jal(im int32) uint32 { return Branch(CondAL, 1, BranchOffset(im)) }
func Ret() uint32         { return Mov(PC, LR) }

// memory operations

func offAbs(off int32) (bitU, uoff uint32) {
	if off >= 0 {
		return 1, uint32(off)
	} else {
		return 0, uint32(-off)
	}
}

func mem(ret, base uint32, off int32, bitL, bitB uint32) uint32 {
	bitU, uoff := offAbs(off)
	return Mem(CondAL, bitL, bitU, bitB, PWOffset, ret, base, AddrImm(uoff))
}

// offset: the absolute value can have 12-bit at maximum

func Ld(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 1, 0) }
func St(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 0, 0) }
func Ldb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 1, 1) }
func Stb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 0, 1) }
