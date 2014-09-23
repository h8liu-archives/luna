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
	w(0xe59f0040, mem(CondAL, 1, 1, 0, pwOffset, r0, pc, addrImm(64)))
	w(0xe3a01001, arith(CondAL, OpMov, 0, 0, r1, shiftIm(1, 0)))
	w(0xe1a01901, arith(CondAL, OpMov, 0, 0, r1, shiftReg(ShSLLImm, r1, 18, 0)))
	w(0xe5801004, mem(CondAL, 0, 1, 0, pwOffset, r1, r0, addrImm(4)))
	w(0xe3a01001, arith(CondAL, OpMov, 0, 0, r1, shiftIm(1, 0)))
	w(0xe1a01801, arith(CondAL, OpMov, 0, 0, r1, shiftReg(ShSLLImm, r1, 16, 0)))

	// loop:
	w(0xe5801028, mem(CondAL, 0, 1, 0, pwOffset, r1, r0, addrImm(40)))
	w(0xe3a0283f, arith(CondAL, OpMov, 0, 0, r2, shiftIm(0x3f, 8)))

	// wait1:
	w(0xe2422001, arith(CondAL, OpSub, 0, r2, r2, shiftIm(1, 0)))
	w(0xe3520000, arith(CondAL, OpCmp, 1, r2, 0, shiftIm(0, 0)))
	w(0x1afffffc, branch(CondNE, 0, branchOffset(-4))) // wait1

	w(0xe580101c, mem(CondAL, 0, 1, 0, pwOffset, r1, r0, addrImm(28)))
	w(0xe3a0283f, arith(CondAL, OpMov, 0, 0, r2, shiftIm(0x3f, 8)))

	// wait2:
	w(0xe2422001, arith(CondAL, OpSub, 0, r2, r2, shiftIm(1, 0)))
	w(0xe3520000, arith(CondAL, OpCmp, 1, r2, 0, shiftIm(0, 0)))
	w(0x1afffffc, branch(CondNE, 0, branchOffset(-4))) // wait2
	w(0xeafffff4, branch(CondAL, 0, branchOffset(-12)))

	// end:
	w(0xeafffffe, branch(CondAL, 0, branchOffset(-2))) // unreachable

	// [data]
	w(0x20200000, 0x20200000)
}
