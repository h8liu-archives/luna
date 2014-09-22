package lunasim

/*
A memory address has 32 bits. Each page has 4096 bytes.
A page offset will take 12 bits. MMU will keep the last 12 bits,
and translate the first 20 bits.

The translate has two stages.
The first stage will translate the first 12 bits.
The second stage will tanslate the second 8 bits.
*/

type mmu struct {
	enabled        bool
	baseAddr0      uint32
	baseAddr1      uint32
	memCtrl        *memCtrl
	asid           int
	width          uint // width for using baseAddr0
	cp15           *cp15
	checkUnaligned bool
}

const (
	addrNbit  = 32
	pageNbit  = 12
	pageSize  = 1 << pageNbit
	indexNbit = addrNbit - pageNbit
	lvl2Nbit  = 8
	lvl2Mask  = (1 << lvl2Nbit) - 1
	lvl1Nbit  = indexNbit - lvl2Nbit
)

func newMMU() *mmu {
	ret := new(mmu)
	ret.setWidth(0)
	return ret
}

func (u *mmu) setWidth(w uint) {
	u.width = w
}

func (u *mmu) translate(vaddr uint32, isWrite bool) uint32 {
	if !u.enabled {
		return vaddr
	}
	return u.walk(vaddr, isWrite)
}

func (u *mmu) loadWord(addr uint32) uint32 {
	return u.memCtrl.loadWord(addr)
}

func (u *mmu) useBase0(vaddr uint32) bool {
	if u.width == 0 {
		return true
	}
	return (vaddr >> (addrNbit - u.width)) == 0
}

// lvl1Addr returns the address of the first stage page table
func (u *mmu) lvl1Addr(vaddr uint32) uint32 {
	ret := (vaddr >> (addrNbit - lvl1Nbit)) * 4
	if u.useBase0(vaddr) {
		ret += u.baseAddr0
	} else {
		ret += u.baseAddr1
	}
	return ret
}

// lvl2Addr returns the address of the second stage page table
func (u *mmu) lvl2Addr(vaddr, table uint32) uint32 {
	index := (vaddr >> pageNbit) & lvl2Mask
	return table + index*4
}

func (u *mmu) walk(vaddr uint32, isWrite bool) uint32 {
	panic("todo")
}
