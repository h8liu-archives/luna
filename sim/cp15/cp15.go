package cp15

type Core interface {
	SetUseMMU(b bool)
	SetCheckUnalign(b bool)
}
