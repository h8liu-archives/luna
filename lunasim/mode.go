package lunasim

import (
	"fmt"
)

// All the modes.
const (
	USR Mode = 0x10
	FIQ Mode = 0x11
	IRQ Mode = 0x12
	SVC Mode = 0x13
	MON Mode = 0x16
	ABT Mode = 0x17
	UND Mode = 0x1b
	SYS Mode = 0x1f
)

// Mode is the mode that the processor is in.
type Mode int

var modeString = map[Mode]string{
	USR: "USR",
	FIQ: "FIQ",
	IRQ: "IRQ",
	SVC: "SVC",
	MON: "MON",
	ABT: "ABT",
	UND: "UND",
	SYS: "SYS",
}

var modeGood = map[Mode]bool{
	USR: true,
	FIQ: true,
	IRQ: true,
	SVC: true,
	ABT: true,
	UND: true,
	SYS: true,
}

/* String returns the string of the mode.
e.g.	Mode(USR).String() == "USR"
		Mode(1).String() == "<mode 1>"
*/
func (m Mode) String() string {
	ret := modeString[m]
	if ret != "" {
		return ret
	}
	return fmt.Sprintf("<mode %d>", m)
}

/* IsGood returns if a mode is good.
e.g. 	IsGood(USR) == true
		IsGood(FIQ) == true
		IsGood(MON) == false
*/
func (m Mode) IsGood() bool {
	return modeGood[m]
}
