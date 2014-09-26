package sim

// physical memory
type PhyMemory struct {
	pages []*PhyPage
}

var _ MemArea = new(PhyPage)

func NewPhyMemory(size uint32) *PhyMemory {
	m := new(PhyMemory)
	if size%pageSize != 0 {
		panic("memory size not aligned to page size")
	}

	m.pages = make([]*PhyPage, size/pageSize)

	return m
}

func pageAddr(paddr uint32) (uint32, uint32) {
	return paddr >> pageNbit, paddr & pageMask
}

func (m *PhyMemory) page(index uint32) *PhyPage {
	p := m.pages[index]
	if p == nil {
		p = NewPhyPage()
		m.pages[index] = p
	}
	return p
}

func (m *PhyMemory) pageOff(paddr uint32) (*PhyPage, uint32) {
	index, off := pageAddr(paddr)
	return m.page(index), off
}

func (m *PhyMemory) ReadU8(paddr uint32) uint32 {
	p, off := m.pageOff(paddr)
	return p.ReadU8(off)
}

func (m *PhyMemory) WriteU8(paddr, v uint32) {
	p, off := m.pageOff(paddr)
	p.WriteU8(off, v)
}

func (m *PhyMemory) ReadU32(paddr uint32) uint32 {
	p, off := m.pageOff(paddr)
	return p.ReadU32(off)
}

func (m *PhyMemory) WriteU32(paddr, v uint32) {
	p, off := m.pageOff(paddr)
	p.WriteU32(off, v)
}
