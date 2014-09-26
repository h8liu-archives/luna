// finger is a strict subset of arm instructions that is similar to misp
// and provides a simple yet usable assembly language for easier understanding
package finger

import (
	ar "github.com/h8liu/luna/arm/arith"
	. "github.com/h8liu/luna/arm/branch"
	. "github.com/h8liu/luna/arm/cond"
	. "github.com/h8liu/luna/arm/mem"
	rn "github.com/h8liu/luna/arm/regname"
)

const (
	R0  = 0
	R1  = 1
	R2  = 2
	R3  = 3
	R4  = 4
	R5  = 5
	R6  = 6
	R7  = 7
	R8  = 8
	R9  = 9
	R10 = 10
	R11 = 11
	R12 = 12
	R13 = 13
	R14 = 14
	R15 = 15

	SP = rn.SP // stack pointer
	LR = rn.LR // link register
	PC = rn.PC // program counter
)

// register only operations

func r3(op, ret, r1, r2 uint32) uint32 {
	return ar.Arith(CondAL, op, 0, r1, ret,
		ar.ShiftReg(ar.ShSLLImm, r2, 0, 0),
	)
}

func r3i(op, ret, r1, im uint32) uint32 {
	return ar.Arith(CondAL, op, 0, r1, ret, ar.ShiftIm(im, 0))
}

func And(ret, r1, r2 uint32) uint32 { return r3(ar.OpAnd, ret, r1, r2) }
func Xor(ret, r1, r2 uint32) uint32 { return r3(ar.OpXor, ret, r1, r2) }
func Sub(ret, r1, r2 uint32) uint32 { return r3(ar.OpSub, ret, r1, r2) }
func Cmp(r1, r2 uint32) uint32      { return r3(ar.OpCmp, 0, r1, r2) }
func Add(ret, r1, r2 uint32) uint32 { return r3(ar.OpAdd, ret, r1, r2) }
func Or(ret, r1, r2 uint32) uint32  { return r3(ar.OpOrr, ret, r1, r2) }
func Mov(ret, r1 uint32) uint32     { return r3(ar.OpMov, ret, 0, r1) }
func Not(ret, r1 uint32) uint32     { return r3i(ar.OpBic, ret, r1, 0) }

func Mul(ret, r1, r2 uint32) uint32 {
	return ar.Mul(CondAL, 0, ret, r1, r2)
}

func Noop() uint32 { return 0 }

// ARMv5 does not have division

// register and 8-bit

func Andi(ret, r1, im uint32) uint32 { return r3i(ar.OpAnd, ret, r1, im) }
func Xori(ret, r1, im uint32) uint32 { return r3i(ar.OpXor, ret, r1, im) }
func Subi(ret, r1, im uint32) uint32 { return r3i(ar.OpSub, ret, r1, im) }
func Cmpi(r1, im uint32) uint32      { return r3i(ar.OpCmp, 0, r1, im) }
func Addi(ret, r1, im uint32) uint32 { return r3i(ar.OpAdd, ret, r1, im) }
func Ori(ret, r1, im uint32) uint32  { return r3i(ar.OpOrr, ret, r1, im) }
func Movi(ret, im uint32) uint32     { return r3i(ar.OpMov, ret, 0, im) }

// register shifting with 5-bit im

func si(op, ret, r1, sh uint32) uint32 {
	return ar.Arith(CondAL, ar.OpMov, 0, 0, ret,
		ar.ShiftReg(op, r1, sh, 0),
	)
}

func s(op, ret, r1, r2 uint32) uint32 {
	return ar.Arith(CondAL, ar.OpMov, 0, 0, ret,
		ar.ShiftReg(op, r1, 0, r2),
	)
}

func Sllv(ret, r1, sh uint32) uint32 { return si(ar.ShSLLImm, ret, r1, sh) }
func Srlv(ret, r1, sh uint32) uint32 { return si(ar.ShSRLImm, ret, r1, sh) }
func Srav(ret, r1, sh uint32) uint32 { return si(ar.ShSRAImm, ret, r1, sh) }
func Sll(ret, r1, r2 uint32) uint32  { return s(ar.ShSLLReg, ret, r1, r2) }
func Srl(ret, r1, r2 uint32) uint32  { return s(ar.ShSRLReg, ret, r1, r2) }
func Sra(ret, r1, r2 uint32) uint32  { return s(ar.ShSRAReg, ret, r1, r2) }

// branch operations

func b(cond, bitL uint32, im int32) uint32 {
	return Branch(cond, bitL, BranchOffset(im))
}

func Beq(im int32) uint32 { return b(CondEQ, 0, im) }
func Bne(im int32) uint32 { return b(CondNE, 0, im) }
func Bge(im int32) uint32 { return b(CondGE, 0, im) }
func Bl(im int32) uint32  { return b(CondLT, 0, im) }
func J(im int32) uint32   { return b(CondAL, 0, im) }

func Jal(im int32) uint32 { return b(CondAL, 1, im) }
func Jr(r uint32) uint32  { return Mov(PC, r) }
func Ret() uint32         { return _ret }

var _ret = Jr(LR)

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
	return Mem(CondAL, bitL, bitU, bitB, PWOffset,
		ret, base, AddrImm(uoff),
	)
}

// offset: the absolute value can have 12-bit at maximum

func Ld(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 1, 0) }
func St(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 0, 0) }
func Ldb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 1, 1) }
func Stb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 0, 1) }
