// Prints "Hello, world." on serial port
package hello

import (
	"bytes"
	"encoding/binary"

	. "github.com/h8liu/luna/finger"
)

// an image with two pages
// code is mapped to 0x8000
// data is mapped to 0x9000
// both code and data must be fix in one page.
func Img() (code, data []byte) {
	cbuf := new(bytes.Buffer)
	b4 := make([]byte, 4)

	for _, b := range []uint32{
		Movi(R1, 9),
		Sllv(R1, R1, 12),
		Ld(R2, R1, 0), // gpio base
		Ld(R3, R1, 4), // uart0 base

		Movi(R0, 0),
		St(R0, R3, 0x30), // disable uart0
		St(R0, R2, 0x94), // disable pull up/down
		// delay150
		Movi(R0, 150),
		Subi(R0, R0, 1),
		Cmpi(R0, 0),
		Bne(-4),

		Movi(R0, 0x3),
		Sllv(R0, R0, 14),
		St(R0, R2, 0x98), // GPPUDCLK0 for pin 14 & 15
		// delay150
		Movi(R0, 150),
		Subi(R0, R0, 1),
		Cmpi(R0, 0),
		Bne(-4),

		Movi(R0, 0),
		St(R0, R2, 0x98), // GPPUDCLK0 make it take effect

		Movi(R0, 0x7),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0xff), // 0x7ff
		St(R0, R3, 0x44),

		Movi(R0, 1),
		St(R0, R3, 0x24), // IBRD
		Movi(R0, 40),
		St(R0, R3, 0x28), // FBRD

		Movi(R0, 0x7),
		Sllv(R0, R0, 4),
		St(R0, R3, 0x2C), // LCRH

		Movi(R0, 0x7),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0xf2),
		St(R0, R3, 0x38), // IMSC

		Movi(R0, 0x3),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0x1),
		St(R0, R3, 0x30), // enable uart

		Addi(R0, R1, 8), // start of the string

		// loop:
		Ldb(R4, R0, 0),
		Cmpi(R4, 0),
		Beq(7), // end

		// byte is now saved in R4
		// wait_loop:
		Ld(R5, R3, 0x18),
		Srlv(R5, R5, 5),
		Andi(R5, R5, 0x1),
		Cmpi(R5, 0),
		Bne(-6), // break on zero, continue pulling on 1

		St(R4, R3, 0), // write the byte out
		Addi(R0, R0, 1),
		J(-12), // goto loop

		// end:
		J(-2), // happily ever after
	} {
		binary.LittleEndian.PutUint32(b4, b)
		cbuf.Write(b4)
	}

	dbuf := new(bytes.Buffer)
	for _, b := range []uint32{
		0x20200000,
		0x20201000,
	} {
		binary.LittleEndian.PutUint32(b4, b)
		dbuf.Write(b4)
	}

	dbuf.Write([]byte("Hello, world.\n"))
	dbuf.Write([]byte{0})

	return cbuf.Bytes(), dbuf.Bytes()
}
