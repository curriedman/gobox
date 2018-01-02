package bounded

type Interface interface {
	FuncMinBound() Interface
	FuncMaxBound() Interface
}
