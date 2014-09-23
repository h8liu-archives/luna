package sim

import (
	"errors"
	"fmt"
)

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

	cpu *CPU
}

const (
	addrNbit  = 32
	pageNbit  = 12
	pageSize  = 1 << pageNbit
	pageMask  = pageSize - 1
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

func (u *mmu) translate(vaddr uint32, isWrite bool) (uint32, error) {
	if !u.enabled {
		return vaddr, nil
	}
	return u.walk(vaddr, isWrite)
}

func (u *mmu) readU32(addr uint32) uint32 {
	return u.memCtrl.readU32(addr)
}

func (u *mmu) useBase0(vaddr uint32) bool {
	if u.width == 0 {
		return true
	}
	return (vaddr >> (addrNbit - u.width)) == 0
}

// lvl1Addr returns the address of the first stage page table
func (u *mmu) lvl1EntryAddr(vaddr uint32) uint32 {
	ret := (vaddr >> (addrNbit - lvl1Nbit)) * 4
	if u.useBase0(vaddr) {
		ret += u.baseAddr0
	} else {
		ret += u.baseAddr1
	}
	return ret
}

// lvl2Addr returns the address of the second stage page table
func (u *mmu) lvl2EntryAddr(vaddr, table uint32) uint32 {
	index := (vaddr >> pageNbit) & lvl2Mask
	return (table & 0xfffffc00) + index*4
}

func errPermFault(ap2 uint32, ap10 uint32) error {
	return fmt.Errorf("perm fault: ap2=%d ap10=%d", ap2, ap10)
}

var errReserved = errors.New("reserved")

// checkPerm checks the permission?
// returns an error
func (u *mmu) checkPerm(vaddr uint32, ap2 uint32, ap10 uint32,
	isWrite bool, isSection bool) error {
	switch ap2 {
	case 1:
		switch ap10 {
		case 0:
			return errReserved
		case 1:
			if isWrite || !u.cpu.isPriviledged() {
				return errPermFault(ap2, ap10)
			}
		case 2:
			// Deprecated // TODO: why?
			if isWrite {
				return errPermFault(ap2, ap10)
			}
		case 3:
			if isWrite {
				if isSection {
					u.cp15.setMemoryAbort(vaddr, permFaultSection, isWrite)
				} else {
					u.cp15.setMemoryAbort(vaddr, permFaultPage, isWrite)
				}
				return errPermFault(ap2, ap10)
			}
		default:
			panic("bug")
		}
	case 0:
		switch ap10 {
		case 0:
			if isSection {
				u.cp15.setMemoryAbort(vaddr, permFaultSection, isWrite)
			} else {
				u.cp15.setMemoryAbort(vaddr, permFaultPage, isWrite)
			}
			return errPermFault(ap2, ap10)
		case 1:
			if !u.cpu.isPriviledged() {
				return errPermFault(ap2, ap10)
			}
		case 2:
			if isWrite && !u.cpu.isPriviledged() {
				return errPermFault(ap2, ap10)
			}
		case 3:
			// do nothing
		default:
			panic("bug")
		}
	default:
		panic("bug")
	}

	return nil
}

func (u *mmu) checkPermLvl1(vaddr uint32, e uint32, isWrite bool) error {
	ap2 := (e >> 15) & 0x1
	ap10 := (e >> 10) & 0x3
	return u.checkPerm(vaddr, ap2, ap10, isWrite, true)
}

func (u *mmu) checkPermLvl2(vaddr uint32, e uint32, isWrite bool) error {
	ap2 := (e >> 9) & 0x1
	ap10 := (e >> 4) & 0x3
	return u.checkPerm(vaddr, ap2, ap10, isWrite, false)
}

func (u *mmu) needPermCheck(entry uint32, isSuper bool) (bool, error) {
	var domIndex int
	if !isSuper {
		domIndex = int((entry >> 5) & 0xf)
	}

	domain := u.cp15.domains[domIndex]

	switch domain {
	case 0:
		return false, errors.New("domain fault")
	case 1:
		return true, nil
	case 2:
		return false, errors.New("domain reserved")
	case 3:
		return false, nil
	}

	panic("bug")
}

func trans(e uint32, vaddr uint32) uint32 {
	return (e & ^uint32(pageMask)) | (vaddr & pageMask)
}

func (u *mmu) walk(vaddr uint32, isWrite bool) (uint32, error) {
	ad1 := u.lvl1EntryAddr(vaddr)
	e1 := u.readU32(ad1)
	switch e1 & 0x3 {
	case 0x0: // unmapped
		u.cp15.setMemoryAbort(vaddr, transFaultSection, isWrite)
	case 0x1: // look into coarse second-level table
		// do nothing
	case 0x2: // section for modified virtual address
		isSuper := (e1 & (0x1 << 18)) == 0
		needCheck, err := u.needPermCheck(e1, isSuper)
		if err != nil {
			return 0, err
		}

		if needCheck {
			if err = u.checkPermLvl1(vaddr, e1, isWrite); err != nil {
				return 0, err
			}
		}

		if isSuper {
			panic("todo: supersection")
		}

		return trans(e1, vaddr), nil
	case 0x3:
		return 0, errors.New("reserved")
	default:
		panic("bug")
	}

	ad2 := u.lvl2EntryAddr(vaddr, e1)
	e2 := u.readU32(ad2)

	needCheck, err := u.needPermCheck(e1, false)
	if err != nil {
		return 0, err
	}
	if needCheck {
		if err = u.checkPermLvl2(vaddr, e2, isWrite); err != nil {
			return 0, err
		}
	}

	switch e2 & 0x3 {
	case 0x0:
		u.cp15.setMemoryAbort(vaddr, transFaultPage, isWrite)
		return 0, errors.New("page fault")
	case 0x1:
		panic("todo: large page")
	case 0x2:
		return trans(e2, vaddr), nil
	case 0x3:
		panic("todo: extended small page")
	default:
		panic("bug")
	}
}
