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

// Currnet program status register
type cpsr struct {
	n, z, c, v, q, e bool
	a, i, f, t, m    bool
}

// Saved program status register
type spsr struct {
	svc cpsr
	mon cpsr
	abt cpsr
	und cpsr
	irq cpsr
	fiq cpsr
}

type CPU struct {
	regs *regs
	cpsr *cpsr
	spsr *spsr
}

func NewCPU() *CPU {
	ret := new(CPU)
	ret.regs = newRegs()
	ret.cpsr = new(cpsr)
	ret.spsr = new(spsr)

	return ret
}

func (cpu *CPU) isPriviledged() bool {
	panic("todo")
}
