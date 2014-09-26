package sim

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

	phyMem *PhyMemory
}

func NewCPU(mem *PhyMemory) *CPU {
	ret := new(CPU)
	ret.regs = newRegs()
	ret.cpsr = new(cpsr)
	ret.spsr = new(spsr)

	ret.phyMem = mem

	return ret
}

func (cpu *CPU) isPriviledged() bool {
	panic("todo")
}

func (cpu *CPU) Reset(pc uint32) {
	// TODO:
}

func (cpu *CPU) Step() {
	// TODO:
}
