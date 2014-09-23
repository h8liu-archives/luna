package asm

const (
	pwMask   = 0x9
	addrMask = (0x1 << 25) | 0xffff
)

const (
	PWOffset   = 0x8 // use Rn+offset, do not write back
	PWPre      = 0x9 // use Rn+offset, and write back
	PWPost     = 0x0 // use Rn, write back Rn+offset
	PWPostUser = 0x1 // use Rn, write back Rn+offset, as if in user mode
)

const (
	AddrSll = 0       // shift right logic
	AddrSrl = 1       // shift right logic
	AddrSra = 4       // shift right arithmetic
	AddrSrr = 6       // shift right rotate
	AddrSrc = 6 + 0xf // shift right carry
)

func AddrImm(im uint32) uint32 {
	return im & 0xfff
}

func AddrReg(mode, im, rm uint32) uint32 {
	ret := uint32(0x1 << 25) // set the I bit first
	ret |= (mode & 0x7) << 4
	if mode != AddrSrc {
		ret |= (im & 0x1f) << 7
	}
	ret |= rm & 0xf
	return ret
}

// rd is the address for store/load
// rn is the base register for calculating adderss
// addr has the bitI and the addr-operand
// bitL tells if it is a load; 1 for load, 0 for store.
// bitU tells if it is a plus or minus on rn; 1 for plus, 0 for minus
// bitB tells if it is a byte load/store operation; 1 for byte, 0 for word
// bitPW tells the accessing mode
func Mem(cond, bitL, bitU, bitB, bitPW, rd, rn, addr uint32) uint32 {
	ret := cond << condShift
	ret |= 0x1 << 26 // memory op
	ret |= (bitL & 0x1) << 20
	ret |= (bitU & 0x1) << 23
	ret |= (bitB & 0x1) << 22
	ret |= (bitPW & pwMask) << 21
	ret |= (rn & 0xf) << 16
	ret |= (rd & 0xf) << 12
	ret |= addr & addrMask
	return ret
}
