package main

import (
	"github.com/h8liu/luna/arm"

	"encoding/binary"
	"fmt"
)

func main() {
	buf := arm.MakeTestBinary()

	n := len(buf)
	if n%4 != 0 {
		panic("buffer not aligned")
	}

	for i := 0; i < n; i += 4 {
		inst := binary.LittleEndian.Uint32(buf[i : i+4])
		fmt.Printf("%08x\n", inst)
	}
}
