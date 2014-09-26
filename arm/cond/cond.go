package cond

const (
	CondEQ = 0   // z set
	CondNE = 1   // z clear
	CondCS = 2   // c set
	CondCC = 3   // c clear
	CondMI = 4   // n set, minus
	CondPL = 5   // n clear, plus
	CondVS = 6   // v set, overflow
	CondVC = 7   // v clear, no overflow
	CondHI = 8   // c set and z clear,
	CondLS = 9   // c clear or z set
	CondGE = 0xA // n == v
	CondLT = 0xB // n != v
	CondGT = 0xC // z == 0, n == v
	CondLE = 0xD // z == 1 or n != v
	CondAL = 0xE // always
	CondIN = 0xF // use for additional instruction coding
)

const (
	CondShift = 28
)

func SetCond(i uint32, cond uint32) uint32 {
	return (i & ^uint32(0xf0000000)) | ((cond & 0xf) << 28)
}

func GetCond(i uint32) uint8 {
	return uint8(i >> 28)
}
