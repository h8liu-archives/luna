package lunasim

type cp15 struct {
	domains []int
}

const (
	transFaultSection = 0x5
	transFaultPage    = 0x7
	permFaultSection  = 0xd
	permFaultPage     = 0xf
)

func (p *cp15) setMemoryAbort(vaddr uint32, status uint32, isWrite bool) {
	panic("todo")
}
