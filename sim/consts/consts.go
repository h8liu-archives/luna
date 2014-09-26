package consts

const (
	AddrNbit  = 32            // memory address width
	PageNbit  = 12            // page offset width
	PageSize  = 1 << PageNbit // page size in bytes
	PageMask  = PageSize - 1  // page offset mask
	IndexNbit = AddrNbit - PageNbit
)
