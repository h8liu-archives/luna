package asm

func branchOffset(i int32) uint32 {
	return uint32(i)
}

func branch(cond, bitL, im24 uint32) uint32 {
	ret := cond << condShift
	ret |= 0x5 << 25
	ret |= (bitL & 0x1) << 24
	ret |= im24 & bitMask24
	return ret
}
