package sim

const Nreg = 16

type regs struct {
	usr []uint32
	svc []uint32
	mon []uint32
	abt []uint32
	und []uint32
	irq []uint32
	fiq []uint32
}

// TODO: not sure if there are those many register
// some doc says only fiq has its own register set
func newRegs() *regs {
	ret := new(regs)

	ret.usr = make([]uint32, Nreg)
	ret.svc = make([]uint32, Nreg)
	ret.mon = make([]uint32, Nreg)
	ret.abt = make([]uint32, Nreg)
	ret.und = make([]uint32, Nreg)
	ret.irq = make([]uint32, Nreg)
	ret.fiq = make([]uint32, Nreg)

	return ret
}
