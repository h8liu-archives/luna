package sim

import (
	"encoding/binary"
)

type MemArea interface {
	ReadU8(offset uint32) uint32
	WriteU8(offset, v uint32)
	ReadU32(offset uint32) uint32
	WriteU32(offset, v uint32)
}

type PhyPage struct {
	dat []byte
}

var _ MemArea = new(PhyPage)

func alignU32(offset uint32) uint32 {
	return offset & ^uint32(0x3)
}

func NewPhyPage() *PhyPage {
	ret := new(PhyPage)
	ret.dat = make([]byte, pageSize)
	return ret
}

var endian = binary.LittleEndian

func (p *PhyPage) ReadU8(offset uint32) uint32 {
	return uint32(p.dat[offset&pageMask])
}

func (p *PhyPage) WriteU8(offset, v uint32) {
	p.dat[offset&pageMask] = uint8(v)
}

func (p *PhyPage) sliceU32(offset uint32) []byte {
	offset &= pageMask
	offset = alignU32(offset)
	return p.dat[offset : offset+4]
}

func (p *PhyPage) ReadU32(offset uint32) uint32 {
	return endian.Uint32(p.sliceU32(offset))
}

func (p *PhyPage) WriteU32(offset, v uint32) {
	endian.PutUint32(p.sliceU32(offset), v)
}
