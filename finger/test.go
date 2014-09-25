package finger

import (
	"bytes"
	"encoding/binary"
)

func TestBin() (code, data []byte) {
	buf := new(bytes.Buffer)
	b4 := make([]byte, 4)

	for _, b := range []uint32{
		// start:
		Movi(SP, 1),      // sp = 1
		Sllv(SP, SP, 12), // sp = sp << 12 // data page offset

		Ld(R0, SP, 0),    // r0 = [sp]
		Movi(R1, 1),      // r1 = 1
		Sllv(R1, R1, 18), // r1 = r1 << 18
		St(R1, R0, 4),    // [r0 + 4] = r1
		Movi(R1, 1),      // r1 = 1
		Sllv(R1, R1, 16), // r1 = r1 << 16
		Movi(R3, 0),      // r3 = 0
		Movi(R4, 0x3f),   // r4 = 0x3f

		// loop:
		St(R1, R0, 40),   // [r0 + 40] = r1
		Sllv(R2, R4, 16), // r2 = r4 << 16

		// wait1:
		Subi(R2, R2, 1), // r2 = r2 - 1
		Cmp(R2, R3),     // if r2 != r3:
		Bne(-4),         // goto wait1

		St(R1, R0, 28),   // [r0 + 28] = r1
		Sllv(R2, R4, 16), // r2 = r4 << 16

		// wait2:
		Subi(R2, R2, 1), // r2 = r2 - 1
		Cmp(R2, R3),     // if r2 != r3:
		Bne(-4),         // goto wait2

		J(-12), // goto loop
	} {
		binary.LittleEndian.PutUint32(b4, b)
		buf.Write(b4)
	}

	dataBuf := new(bytes.Buffer)
	binary.LittleEndian.PutUint32(b4, 0x20200000)

	return buf.Bytes(), dataBuf.Bytes()
}
