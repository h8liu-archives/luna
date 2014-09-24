// finger is a strict subset of arm instructions that is similar to misp
// and provides a simple yet usable assembly language for easier understanding
package finger

import (
	"github.com/h8liu/luna/arm"
)

// register only operations

func r3(op, ret, r1, r2 uint32) uint32 {
	return arm.Arith(arm.CondAL, op, 0, r1, ret,
		arm.ShiftReg(arm.ShSLLImm, r2, 0, 0),
	)
}

func r3i(op, ret, r1, im uint32) uint32 {
	return arm.Arith(arm.CondAL, op, 0, r1, ret, arm.ShiftIm(im, 0))
}

func And(ret, r1, r2 uint32) uint32 { return r3(arm.OpAnd, ret, r1, r2) }
func Xor(ret, r1, r2 uint32) uint32 { return r3(arm.OpXor, ret, r1, r2) }
func Sub(ret, r1, r2 uint32) uint32 { return r3(arm.OpSub, ret, r1, r2) }
func Cmp(r1, r2 uint32) uint32      { return r3(arm.OpCmp, 0, r1, r2) }
func Add(ret, r1, r2 uint32) uint32 { return r3(arm.OpAdd, ret, r1, r2) }
func Or(ret, r1, r2 uint32) uint32  { return r3(arm.OpOrr, ret, r1, r2) }
func Mov(ret, r1 uint32) uint32     { return r3(arm.OpMov, ret, 0, r1) }
func Not(ret, r1 uint32) uint32     { return r3i(arm.OpBic, ret, r1, 0) }

func Mul(ret, r1, r2 uint32) uint32 {
	return arm.Mul(arm.CondAL, 0, ret, r1, r2)
}

// ARMv5 does not have division

// register and 8-bit

func Andi(ret, r1, im uint32) uint32 { return r3i(arm.OpAnd, ret, r1, im) }
func Xori(ret, r1, im uint32) uint32 { return r3i(arm.OpXor, ret, r1, im) }
func Subi(ret, r1, im uint32) uint32 { return r3i(arm.OpSub, ret, r1, im) }
func Addi(ret, r1, im uint32) uint32 { return r3i(arm.OpAdd, ret, r1, im) }
func Ori(ret, r1, im uint32) uint32  { return r3i(arm.OpOrr, ret, r1, im) }
func Movi(ret, im uint32) uint32     { return r3i(arm.OpMov, ret, 0, im) }

// register shifting with 5-bit im

func si(op, ret, r1, sh uint32) uint32 {
	return arm.Arith(arm.CondAL, arm.OpMov, 0, 0, ret,
		arm.ShiftReg(op, r1, sh, 0),
	)
}

func s(op, ret, r1, r2 uint32) uint32 {
	return arm.Arith(arm.CondAL, arm.OpMov, 0, 0, ret,
		arm.ShiftReg(op, r1, 0, r2),
	)
}

func Sllv(ret, r1, sh uint32) uint32 { return si(arm.ShSLLImm, ret, r1, sh) }
func Srlv(ret, r1, sh uint32) uint32 { return si(arm.ShSRLImm, ret, r1, sh) }
func Srav(ret, r1, sh uint32) uint32 { return si(arm.ShSRAImm, ret, r1, sh) }
func Sll(ret, r1, r2 uint32) uint32  { return s(arm.ShSLLReg, ret, r1, r2) }
func Srl(ret, r1, r2 uint32) uint32  { return s(arm.ShSRLReg, ret, r1, r2) }
func Sra(ret, r1, r2 uint32) uint32  { return s(arm.ShSRAReg, ret, r1, r2) }

// branch operations

func b(cond, bitL uint32, im int32) uint32 {
	return arm.Branch(cond, bitL, arm.BranchOffset(im))
}

func Beq(im int32) uint32 { return b(arm.CondEQ, 0, im) }
func Bne(im int32) uint32 { return b(arm.CondNE, 0, im) }
func Bge(im int32) uint32 { return b(arm.CondMI, 0, im) }
func Bl(im int32) uint32  { return b(arm.CondPL, 0, im) }
func J(im int32) uint32   { return b(arm.CondAL, 0, im) }

func Jal(im int32) uint32 { return b(arm.CondAL, 1, im) }
func Jr(r uint32) uint32  { return Mov(arm.PC, r) }
func Ret() uint32         { return Jr(arm.LR) }

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
	return arm.Mem(arm.CondAL, bitL, bitU, bitB, arm.PWOffset,
		ret, base, arm.AddrImm(uoff),
	)
}

// offset: the absolute value can have 12-bit at maximum

func Ld(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 1, 0) }
func St(ret, b uint32, off int32) uint32  { return mem(ret, b, off, 0, 0) }
func Ldb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 1, 1) }
func Stb(ret, b uint32, off int32) uint32 { return mem(ret, b, off, 0, 1) }
