// Physical page
package phypage

import (
	"encoding/binary"

	. "github.com/h8liu/luna/sim/consts"
)

type PhyPage struct {
	dat []byte
}

func alignU32(offset uint32) uint32 {
	return offset & ^uint32(0x3)
}

func NewPhyPage() *PhyPage {
	ret := new(PhyPage)
	ret.dat = make([]byte, PageSize)
	return ret
}

// Raspberry Pi uses little endian
var endian = binary.LittleEndian

func (p *PhyPage) ReadU8(offset uint32) uint32 {
	return uint32(p.dat[offset&PageMask])
}

func (p *PhyPage) WriteU8(offset, v uint32) {
	p.dat[offset&PageMask] = uint8(v)
}

func (p *PhyPage) sliceU32(offset uint32) []byte {
	offset &= PageMask
	offset = alignU32(offset)
	return p.dat[offset : offset+4]
}

func (p *PhyPage) ReadU32(offset uint32) uint32 {
	return endian.Uint32(p.sliceU32(offset))
}

func (p *PhyPage) WriteU32(offset, v uint32) {
	endian.PutUint32(p.sliceU32(offset), v)
}
