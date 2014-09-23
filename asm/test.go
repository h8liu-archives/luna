package asm

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func MainTest() {
	buf := new(bytes.Buffer)
	b := make([]byte, 4)
	w := func(expect, i uint32) {
		fmt.Printf("%08x\n", i)
		binary.LittleEndian.PutUint32(b, i)
		buf.Write(b)
		if i != expect {
			panic(fmt.Sprintf("expect %08x, got %08x\n",
				expect, i))
		}
	}

	// [code]
	w(0xe59f0040, Mem(CondAL, 1, 1, 0, PWOffset, R0, PC, AddrImm(64)))
	w(0xe3a01001, Arith(CondAL, OpMov, 0, 0, R1, ShiftIm(1, 0)))
	w(0xe1a01901, Arith(CondAL, OpMov, 0, 0, R1, ShiftReg(ShSLLImm, R1, 18, 0)))
	w(0xe5801004, Mem(CondAL, 0, 1, 0, PWOffset, R1, R0, AddrImm(4)))
	w(0xe3a01001, Arith(CondAL, OpMov, 0, 0, R1, ShiftIm(1, 0)))
	w(0xe1a01801, Arith(CondAL, OpMov, 0, 0, R1, ShiftReg(ShSLLImm, R1, 16, 0)))

	// loop:
	w(0xe5801028, Mem(CondAL, 0, 1, 0, PWOffset, R1, R0, AddrImm(40)))
	w(0xe3a0283f, Arith(CondAL, OpMov, 0, 0, R2, ShiftIm(0x3f, 8)))

	// wait1:
	w(0xe2422001, Arith(CondAL, OpSub, 0, R2, R2, ShiftIm(1, 0)))
	w(0xe3520000, Arith(CondAL, OpCmp, 1, R2, 0, ShiftIm(0, 0)))
	w(0x1afffffc, Branch(CondNE, 0, BranchOffset(-4))) // wait1

	w(0xe580101c, Mem(CondAL, 0, 1, 0, PWOffset, R1, R0, AddrImm(28)))
	w(0xe3a0283f, Arith(CondAL, OpMov, 0, 0, R2, ShiftIm(0x3f, 8)))

	// wait2:
	w(0xe2422001, Arith(CondAL, OpSub, 0, R2, R2, ShiftIm(1, 0)))
	w(0xe3520000, Arith(CondAL, OpCmp, 1, R2, 0, ShiftIm(0, 0)))
	w(0x1afffffc, Branch(CondNE, 0, BranchOffset(-4))) // wait2
	w(0xeafffff4, Branch(CondAL, 0, BranchOffset(-12)))

	// end:
	w(0xeafffffe, Branch(CondAL, 0, BranchOffset(-2))) // unreachable

	// [data]
	w(0x20200000, 0x20200000)
}
