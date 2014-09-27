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
		// CR = 0, diable UART
		St(R0, R3, 0x30), // CR, disable uart0

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

		// UART ICR = 0x7ff;
		// Clearing pending interrupts
		// TODO: according to the document, it should write 0x7f2
		// since the 3 bits are not supported and should write zero
		//
		// also, according to the doc, after disable the UART, one should
		// - wait for the end of tx or rx of the current character (how? just spin?)
		// - flush the transmit FIFO by setting the FEN bit to 0 in the LCRH
		// and then one can reprogram stuff, and enable UART.
		Movi(R0, 0x7),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0xff), // 0x7ff
		St(R0, R3, 0x44),

		// IBRD = 1, FBRD = 40
		// TODO: what is the baud rate?
		Movi(R0, 1),
		St(R0, R3, 0x24), // IBRD
		Movi(R0, 40),
		St(R0, R3, 0x28), // FBRD

		// LDRH = 0x70
		// word length = 8 bits
		// FIFO enabled
		// stick parity disabled
		// no two stop bits
		// no parity
		// no send break
		Movi(R0, 0x70),
		St(R0, R3, 0x2C), // LCRH

		// UART IMSC = 0x7f2
		// interrupt mask
		// this will enable all the supported interrupts
		Movi(R0, 0x7),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0xf2),
		St(R0, R3, 0x38), // IMSC

		// CR = 0x301
		// no CTS control enable
		// no RTS control enable
		// set RTS to 0 (since it is not connected, so we don't care)
		// enable rx (bit 9)
		// enable tx (bit 8)
		// loopback disabled (bit 7)
		// uart enabled (bit 0)
		Movi(R0, 0x3),
		Sllv(R0, R0, 8),
		Ori(R0, R0, 0x1),
		St(R0, R3, 0x30), // CR, enable uart

		Addi(R0, R1, 8), // start of the string

		// loop:
		Ldb(R4, R0, 0),
		Cmpi(R4, 0),
		Beq(7), // end

		// byte is now saved in R4
		// wait_loop:
		// read in FR, flag register
		Ld(R5, R3, 0x18),
		// check bit 5, see if tx fifo is full
		Srlv(R5, R5, 5),
		Andi(R5, R5, 0x1),
		Cmpi(R5, 0),
		Bne(-6), // break on zero, continue pulling on 1

		// write DR, the data register
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
