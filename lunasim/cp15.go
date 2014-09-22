package lunasim

/*
cp15is the system control coprocessor. It controls all of the standard memory
and system facilities.
 
*/

const ncp15Reg = 16

type cp15 struct {
	domains []int
	regs []uint32
	cpu *CPU
}

const (
	transFaultSection = 0x5
	transFaultPage    = 0x7
	permFaultSection  = 0xd
	permFaultPage     = 0xf
)

const (
	cpIdCodes = iota
	cpSystemConfig
	cpPageTableCtrl
	cpDomainAccessCtrl
	_
	cpFaultStatus
	cpFaultAddress
	cpCacheWriteBufCtrl
	cpTLBCtrl
	cpCacheLockdown
	cpTLBLockdown
	cpDMACtrl
	_
	cpProcessID
	_
	_
)

func init() {
	if cpProcessID != 13 {
		panic("bug")
	}
}

func (p *cp15) setMemoryAbort(vaddr, status uint32, isWrite bool) {
	panic("todo")
}

func newCp15(cpu *CPU) *cp15 {
	ret := new(cp15)
	ret.regs = make([]uint32, ncp15Reg)
	ret.cpu = cpu

	return ret
}


