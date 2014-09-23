package arm

func BranchOffset(i int32) uint32 {
	return uint32(i)
}

func Branch(cond, bitL, im24 uint32) uint32 {
	ret := cond << condShift
	ret |= 0x5 << 25
	ret |= (bitL & 0x1) << 24
	ret |= im24 & bitMask24
	return ret
}
