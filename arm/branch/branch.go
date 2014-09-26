package branch

import (
	. "github.com/h8liu/luna/arm/cond"
)

const bitMask24 = uint32((0x1 << 24) - 1)

func BranchOffset(i int32) uint32 {
	return uint32(i)
}

func Branch(cond, bitL, im24 uint32) uint32 {
	ret := cond << CondShift
	ret |= 0x5 << 25
	ret |= (bitL & 0x1) << 24
	ret |= im24 & bitMask24
	return ret
}
