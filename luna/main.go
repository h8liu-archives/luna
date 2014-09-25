package main

import (
	"github.com/h8liu/luna/finger"

	// "encoding/binary"
	"fmt"
)

func main() {
	buf := finger.TestBin()

	n := len(buf)
	if n%4 != 0 {
		panic("buffer not aligned")
	}

	fmt.Printf("hello\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%02x ", buf[i])
	}
	/*
		for i := 0; i < n; i += 4 {
			inst := binary.LittleEndian.Uint32(buf[i : i+4])
			fmt.Printf("%08x\n", inst)
		}
	*/
}
